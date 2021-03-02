package db

import (
	"Go/filestore/first/db/mysql"
	"encoding/json"
	"fmt"
)
type User struct {
	Username     string
	Email        string
	Phone        string
	SignupAt     string
	LastActiveAt string
	Status       int
}
func UserSignup(userName string, password string)bool{
	var ok = true
	mysql.Init("Tb_User", false)
	db := mysql.Conn()
	tb_user := mysql.Tb_User{
		User_name:       userName,
		User_pwd:        password,
	}
	if err := db.Create(&tb_user).Error;err != nil{
		fmt.Println("insert failed ,err:", err)
		ok = false
	}
	db.Close()
	return ok
}
//登陆 判断密码是否一致
func UserSignin(username string, encpwd string) bool {
	var ok = true
	mysql.Init("Tb_User", false)
	db := mysql.Conn()
	var user mysql.Tb_User
	err := db.First(&user, "User_name=?" , username).Error
	if err != nil {
		fmt.Println("query failed ,err:", err)
		ok = false
	}
	if user.User_pwd != encpwd{
		fmt.Println("your password or username is incorrect")
		ok = false
	}
	db.Close()
	return ok
}
//更新token，每次登陆token不一样
func UpdateToken(username string, token string) bool {
	var ok = true
	mysql.Init("Tb_User_Token", false)
	db := mysql.Conn()
	err := db.Model(mysql.Tb_User_Token{}).Where("User_name=?", username).Update("User_Token", token).Error
	if err != nil {
		fmt.Println("Update failed, err:", err)
		ok = false
	}
	return ok
}

//: 查询用户信息
func GetUserInfo(username string) (User, bool) {
	var ok = true
	mysql.Init("Tb_User", false)
	db := mysql.Conn()
	var user User
	var tb_user mysql.Tb_User
	err := db.First(&tb_user, "User_name=?", username).Error
	if err != nil {
		fmt.Println("query failed, err:", err)
		ok = false
		return user, ok
	}
	user.Username = tb_user.User_name
	t, _  := json.Marshal(tb_user.CreatedAt)
	user.SignupAt = string(t)
	return user, ok
}
