package db

import "time"

type FileMeta struct {
	//FileSha1 一个类似md5的记录文件信息的
	FileSha1 string
	FileSize int64
	FileName string
	Location string
	UploadAt time.Time
}
var fileMeta map[string]FileMeta
func Init(){
	fileMeta = make(map[string]FileMeta)
}

func GetFileMeta(filename string) FileMeta {
	return fileMeta[filename]
}
func SetFileMeta(fileMeta1 FileMeta){
	fileMeta[fileMeta1.FileSha1] = fileMeta1
}