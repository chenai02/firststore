package main

import (
	"Go/filestore/first/handle"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload",handle.Upload)
	http.HandleFunc("/file/upload/suc", handle.UploadHandle)
	http.HandleFunc("/file/meta", handle.GetFileMetaHeader)
	http.HandleFunc("/user/signup", handle.SignupHandle)
	http.HandleFunc("/user/signup/suc", handle.Signup)
	http.HandleFunc("/user/signin", handle.SigninHandle)
	http.HandleFunc("/user/signin/suc", handle.Signin)
	http.HandleFunc("/static/view/home.html", handle.HomeHandle)
	http.HandleFunc("/user/info", handle.HTTPInterceptor(handle.UserInfoHandle))
	http.ListenAndServe(":8080", nil)
}
