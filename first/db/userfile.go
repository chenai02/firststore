package db

import (
	"Go/filestore/first/db/mysql"
	"fmt"
	"time"
)

func OnupLoadFile(username, filehash, filename string, filesize int64, create time.Time)bool{
	var ok = true
	mysql.Init("Tb_User_File", true)
	db := mysql.Conn()
	user_file := mysql.Tb_User_File{
		User_name: username,
		File_sha1: filehash,
		File_name: filename,
		File_size: filesize,
		CreatedAt: create,
	}
	err := db.Create(&user_file).Error
	if err != nil {
		fmt.Println("create tb_user_file failed, err:",err)
		ok = false
	}
	return ok
}
