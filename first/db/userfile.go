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

func QueryUserFiles(username string, limit int)([]mysql.UserFile, bool){
	var ok = true
	mysql.Init("Tb_User_File", true)
	db := mysql.Conn()
	var user_files []mysql.UserFile
	var tb_user_file []mysql.Tb_User_File
	err := db.Find(&tb_user_file, "User_name=?", username).Error
	if err != nil {
		fmt.Println("get Tb_User_File message failed, err:", err)
		ok = false
		return user_files, ok
	}else if len(tb_user_file) == 0{
		fmt.Printf("the %v is not in db,  \n", username)
	}else {
		if limit > len(tb_user_file) {
			limit = len(tb_user_file)
		}
		for i := 0; i < limit; i++ {
			user_file := mysql.UserFile{
				UserName:    tb_user_file[i].User_name,
				FileHash:    tb_user_file[i].File_sha1,
				FileName:    tb_user_file[i].File_name,
				FileSize:    tb_user_file[i].File_size,
				UploadAt:    tb_user_file[i].CreatedAt,
				LastUpdated: tb_user_file[i].UpdatedAt,
			}
			user_files = append(user_files, user_file)
		}
	}
	return user_files, ok
}
