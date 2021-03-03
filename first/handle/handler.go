package handle

import (
	"Go/filestore/first/db"
	"Go/filestore/first/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func Upload(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		//返回html页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internel server error")
			return
		}
		r.ParseForm()
		username := r.Form.Get("username")
		//fmt.Println("currentUser1:", username)
		SetCurrentUser(username)
		io.WriteString(w, string(data))
	}else if r.Method == "POST" {
		//返回POST页面
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("FormFile failed, err:%v", err)
			return
		}
		defer file.Close()
		//D:\GoPath\src\Go\filestore\first
		//filename := strings.Split(head.Filename, "\\")
		fileMeta := db.FileMeta{
			FileName:head.Filename,
			//Location:"d:/GoPath/src/Go/filestore/first/"+filename[len(filename)-1],
			Location:"d:/GoPath/src/Go/filestore/first/"+head.Filename,
			UploadAt:time.Now(),
		}
		//fmt.Println("FileName:", filename[len(filename)-1])
		dst ,err:= os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("create file failed, err:%v", err)
			return
		}
		defer dst.Close()
		fileMeta.FileSize, err = io.Copy(dst, file)
		//_, err = io.Copy(dst, file)
		if err != nil {
			fmt.Printf("copy file failed, err:%v", err)
			return
		}
		dst.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(dst)
		//meta.SetFileMeta(fileMeta)
		db.SetFileDB(fileMeta)
		//更新用户文件表记录
		//因为js本身问题，导致无法获取当前表单信息，这样做有一个缺陷，无法开多个窗口
		username := GetCurrentUser()
		//r.ParseForm()
		//username := r.Form.Get("username")
		//fmt.Println("currentUser2:", username)
		ok := db.OnupLoadFile(username, fileMeta.FileSha1, fileMeta.FileName, fileMeta.FileSize, fileMeta.UploadAt)
		DeleteCurrentUser()
		if ok {
			http.Redirect(w, r, "/static/view/home.html", http.StatusFound)
		}else{
			w.Write([]byte("Upload Failed."))
		}
	}
}

func UploadHandle(w http.ResponseWriter, r *http.Request){

}

func UploadSave(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "upload success")
}

func GetFileMetaHeader(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost {
		//fileHash := r.Form["fileHash"][0]
		//fileMeta := meta.GetFileMeta(filName)
		//fileMeta := db.GetFileDb(fileHash)
		r.ParseForm()
		username := r.Form.Get("username")
		user_file, ok := db.QueryUserFiles(username, 10)
		if !ok {
			fmt.Println("get userFile failed")
			return
		}
		data, err := json.Marshal(user_file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}
}

