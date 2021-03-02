package handle

import (
	"Go/filestore/first/db"
	"Go/filestore/first/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)
const(
	pass_salt = "*#&%890"
)
//注册接口
func SignupHandle(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}
}
//注册
func Signup(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if len(username) != 0 && (len(username) < 3 || len(password) < 5) {
			w.Write([]byte("Invalid param"))
			return
		}
		if len(username) != 0 {
			//加密密码
			enc_pass := util.Sha1([]byte(password + pass_salt))
			ok := db.UserSignup(username, enc_pass)
			if ok {
				w.Write([]byte("SUCCESS"))
				return
			} else {
				w.Write([]byte("FAILED"))
				return
			}
		}
	}
}
//登陆接口
func SigninHandle(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signin.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}
}
//登陆
func Signin(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		//因为js本身问题，导致无法获取当前表单信息，这样做有一个缺陷，无法开多个窗口
		SetCurrentUser(username)
		//先将密码用相同的方法加密，确保安全性和正确性
		if len(username) != 0 && len(password) != 0 {
			enc_pass := util.Sha1([]byte(password + pass_salt))
			//1.验证用户名及密码
			ok := db.UserSignin(username, enc_pass)
			if !ok {
				w.Write([]byte("FAILED"))
				return
			}
			//2.生成Token
			token := GetToken(username)
			ok = db.UpdateToken(username, token)
			if !ok {
				w.Write([]byte("FAILED"))
				return
			}
			//因为js的问题，导致无法获取currentUser，所以将登陆用户
			//3.登陆后重定向到首页
			//w.Write([]byte("http://" + r.Host + "/static/view/home.html")) 这个不行
			resp := util.RespMsg{
				Code: 0,
				Msg:  "OK",
				Data: struct {
					Location string
					Username string
					Token    string
				}{
					Location: "http://" + r.Host + "/static/view/home.html",
					Username: username,
					Token:    token,
				},
			}
			w.Write(resp.JSONBytes())
		}
	}
}

//显示主页
func HomeHandle(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/home.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}
}

func UserInfoHandle(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost {
		//1.解析请求参数
		r.ParseForm()
		username := r.Form.Get("username")
		//token := r.Form.Get("token")
		//fmt.Println(username)
		////2.验证token
		//isvalidtoken := IsValidToken(token)
		//if !isvalidtoken {
		//	fmt.Println("token time out")
		//	return
		//}
		//3.查询用户
		user, ok := db.GetUserInfo(username)
		if !ok {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		//4.写入body
		//fmt.Println(user.Username, user.SignupAt)
		resp := util.RespMsg{
			Code: 0,
			Msg:  "OK",
			Data: user,
		}
		w.Write(resp.JSONBytes())
	}

}

func GetToken(userName string)string{
	//40位字符：md5(userName+timestamp+token_salt)+timestamp[:8] md5为32位
	ts := fmt.Sprintf("%x", time.Now().Unix())
	return util.MD5([]byte(userName + ts + "_tokensalt"))+ ts[:8]
}

func IsValidToken(token string)bool{
	//验证token的时效性，token后八位为时间，可以对比一下
	return true
}
