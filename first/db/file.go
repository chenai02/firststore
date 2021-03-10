package db

import (
	"Go/filestore/first/db/mysql"
	"fmt"
)

//保存文件信息到数据库中
func OnfileUpload(filehash, filename, fileaddr string, filesize int64) bool {
	mysql.Init("Tb_file", false)
	db := mysql.Conn()
	file_1 := mysql.Tb_File{
		File_sha1: filehash,
		File_size: filesize,
		File_addr: fileaddr,
		File_name: filename,
	}
	//var file mysql.Tb_File
	//db.First(&file, "File_name=?", filename)
	db_err := db.Create(&file_1)
	err := db_err.Error
	if err != nil {
		fmt.Println("autoMigrate failed, err: ", err)
		return false
	}
	//db_err = db.Create(&file_1) //直接报错，因为unique保证数据唯一性
	mysql.Close()
	return true
}

//func main(){
//	ok := OnfileUpload("111", "111", "aa",1)
//	fmt.Println(ok)
//}
func SetFileDB(fileMeta FileMeta) {
	mysql.Init("Tb_File", true)
	db := mysql.Conn()
	tb_file := mysql.Tb_File{
		File_sha1: fileMeta.FileSha1,
		File_name: fileMeta.FileName,
		File_size: fileMeta.FileSize,
		File_addr: fileMeta.Location,
		CreatedAt: fileMeta.UploadAt,
	}
	err := db.Create(&tb_file).Error
	if err != nil {
		fmt.Println("Create fileMeta failed, err:", err)
		return
	}
}
func GetFileDb(fileHash string) mysql.Tb_File {
	var tb_file mysql.Tb_File
	mysql.Init("Tb_file", false)
	db := mysql.Conn()
	db_err := db.First(&tb_file, "File_sha1=?", fileHash)
	if db_err.Error != nil {
		return mysql.Tb_File{}
	}
	return tb_file
	//filemeta := meta.FileMeta{
	//	FileSha1: tb_file.File_sha1,
	//	FileSize: tb_file.File_size,
	//	FileName: tb_file.File_name,
	//	Location: tb_file.File_addr,
	//	UploadAt: strconv.Itoa(tb_file.UpdatedAt),
	//}
	//return filemeta,nil
}
