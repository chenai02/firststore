package handle

import (
	"Go/filestore/first/cache/redis"
	"Go/filestore/first/util"
	"fmt"
	"math"
	"net/http"
	"strconv"
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