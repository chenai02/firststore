package mysql

import (
	"fmt"
	//"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var (
	db *gorm.DB
	err error
	)
func Init(tb string, ok bool)  {
	db, err = gorm.Open("mysql","root:123456@(127.0.0.1:3306)/fileserver?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("connect db failed, err:",err)
		return
	}
	if ok && tb == "Tb_File"{
		db_err := db.Set("db", "ENGINE=InnoDB").AutoMigrate(&Tb_File{})
		if err = db_err.Error; err != nil {
			fmt.Println("autoMigrate failed, err: ", err)
			return
		}
	}
	if ok && tb == "Tb_User"{
		db_err := db.Set("db", "ENGINE=InnoDB").AutoMigrate(&Tb_User{})
		if err = db_err.Error; err != nil {
			fmt.Println("autoMigrate failed, err: ", err)
			return
		}
	}
	if ok && tb == "Tb_User_Token"{
		db_err := db.Set("db", "ENGINE=InnoDB").AutoMigrate(&Tb_User_Token{})
		if err = db_err.Error; err != nil {
			fmt.Println("autoMigrate failed, err: ", err)
			return
		}
	}
	if ok && tb == "Tb_User_File"{
		db_err := db.Set("db", "ENGINE=InnoDB").AutoMigrate(&Tb_User_File{})
		if err = db_err.Error; err != nil {
			fmt.Println("autoMigrate failed, err: ", err)
			return
		}
	}


}

func Conn()*gorm.DB{
	return db
}
func Close()bool{
	err = db.Close()
	if err != nil {
		fmt.Println("close db failed, err:", err)
		return false
	}
	return true
}