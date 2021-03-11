package handle

import (
	"Go/filestore/first/cache/redis"
	"Go/filestore/first/db"
	Redis"github.com/garyburd/redigo/redis"
	"Go/filestore/first/util"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type  MultipartUploadInfo struct {
	FileHash string
	FileSize int
	UploadID string
	ChunkSzie int
	ChunkCount int
}

func InitialMultipartUploadInfohandle(w http.ResponseWriter, r *http.Request)  {

	//1.解析请求参数
	r.ParseForm()
	username := r.Form.Get("username")
	filehash := r.Form.Get("filehash")
	filesize,err := strconv.Atoi(r.Form.Get("filesize"))
	if err != nil {
		w.Write(util.NewRespMsg(-1, "params invalid", nil).JSONBytes())
		return
	}
	//2.获取一个redis连接
	conn := redis.RedisPool().Get()
	defer conn.Close()
	//3.生成分块上传的初始化信息
	upinfo := MultipartUploadInfo{
		FileHash:   filehash,
		FileSize:   filesize,
		UploadID:   username + fmt.Sprintf("%x", time.Now().UnixNano()),
		ChunkSzie:  5 * 1024 * 1024, //5MB
		ChunkCount: int(math.Ceil(float64(filesize)/(5 * 1024 * 1024))),
	}
	//4.将初始化信息写入redis连接池中
	conn.Do("HSET", "MP_" + upinfo.UploadID, "filehash", upinfo.FileHash)
	conn.Do("HSET", "MP_" + upinfo.UploadID, "filesize", upinfo.ChunkSzie)
	conn.Do("HSET", "MP_" + upinfo.UploadID, "chunkcount", upinfo.ChunkCount)
	//5.将响应初始化数据返回到客户端
	w.Write(util.NewRespMsg(0, "OK", upinfo).JSONBytes())
	
}

func UploadPartHandle(w http.ResponseWriter, r *http.Request){

	//1.获取请求参数
	r.ParseForm()
//	username := r.Form.Get("username")
	uploadid := r.Form.Get("uploadid")
	chunkindex := r.Form.Get("index")
	//2.创建一个redis连接用于存储进度
	conn := redis.RedisPool().Get()
	defer conn.Close()
	//3.创建一个文件句柄用以接收分块文件
	fd, err := os.Create("d:/GoPath/src/Go/filestore/first/tmp/"+ uploadid+"/"+chunkindex)
	if err != nil {
		w.Write(util.NewRespMsg(-1, "Upload part failed", nil).JSONBytes())
		return
	}
	defer fd.Close()
	buf := make([]byte, 1024 * 1024)
	for{
		n, err := r.Body.Read(buf)
		fd.Write(buf[:n])
		if err != nil {
			break
		}
	}
	//4.更新redis状态  1代表这一部分完成上传
	conn.Do("HSET", "MP_"+ uploadid, "chunkindex_" + chunkindex, 1)
	//5.返回客户端
	w.Write(util.NewRespMsg(0, "OK", nil).JSONBytes())
}


func CompleteUploadHandle(w http.ResponseWriter, r *http.Request){

	//1.解析参数
	r.ParseForm()
	upid := r.Form.Get("uploadid")
	username := r.Form.Get("username")
	filehash := r.Form.Get("filehash")
	filesize := r.Form.Get("filesize")
	filename := r.Form.Get("filename")
	//2.创建redis连接
	conn := redis.RedisPool().Get()
	defer conn.Close()
	//3.判断分块上传是否全部完成
	//Values将一个空接口类型的数据转化为切片类型的空接口数据
	data, err := Redis.Values(conn.Do("HGETALL", "MP_" + upid))
	if err != nil {
		w.Write(util.NewRespMsg(-1, "complete upload failed", nil).JSONBytes())
		return
	}
	totalCount := 0
	chunkCount := 0
	for i := 0; i < len(data); i += 2{
		k := string(data[i].([]byte))
		v := string(data[i+1].([]byte))
		if k == "chunkcount"{
			totalCount, _ = strconv.Atoi(v)
		}else if strings.HasPrefix(k, "chunkindex_") && v == "1"{
				chunkCount ++
		}
	}
	if totalCount != chunkCount{
		w.Write(util.NewRespMsg(-2, "invalid request", nil).JSONBytes())
		return
	}
	//4.合并分块

	//5.更新唯一文件表和用户文件表
	fsize, _:= strconv.Atoi(filesize)
	db.OnfileUpload(filehash, filename, "", int64(fsize))
	db.OnupLoadFile(username, filehash, filename, int64(fsize), time.Now())
	//6.返回给客户端
	w.Write(util.NewRespMsg(0, "OK", nil).JSONBytes())
}
