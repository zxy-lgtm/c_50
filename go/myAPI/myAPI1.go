package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//储存用户信息的结构体和映射
type User struct {
	Name     string `json : "name"`
	Phone    string `json : "phone"`
	Password string `json : "password"`
}

var userData = make(map[string]User)

//注册函数
func Signin(w http.ResponseWriter, r *http.Request) {
	var user User
	fmt.Fprintln(w, "你好!新用户!\n注册中......")

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&user)
	if _, ok := userData[user.Name]; !ok {
		fmt.Fprintln(w, "注册成功")
		userData[user.Name] = user
	} else {
		fmt.Fprintln(w, "注册失败,该用户已存在")
	}
	//fmt.Println(user.Name)
}

//登录结构体
type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json: "password"`
}

//登录函数
func Login(w http.ResponseWriter, r *http.Request) {
	new := new(LoginRequest)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(new)

	myuser, ok := userData[new.Name]
	if !ok {
		fmt.Fprintln(w, "该用户不存在!!")
		return
	}

	if new.Password != myuser.Password {
		fmt.Fprintln(w, "密码输入错误!!")
		return
	}

	fmt.Fprintln(w, "登陆成功!")

}

type AmendNameRequest struct {
	OldName  string `json : "oldname"`
	NewName  string `json : "newname"`
	Password string `json : "password"`
}

func AmendName(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "你好!用户!\n修改中......")
	var request AmendNameRequest

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&request)

	myuser, ok := userData[request.OldName]
	if !ok {
		fmt.Fprintln(w, "该用户不存在!!")
		return
	}

	if myuser.Password != request.Password {
		fmt.Fprintln(w, "密码输入错误!!")
		return
	}

	myuser.Name = request.NewName
	userData[myuser.Name] = myuser

	fmt.Fprintln(w, "修改成功!")
}

//修改密码结构体
type AmendPasswordRequest struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

//修改密码函数
func AmendPassword(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "你好!用户!\n修改中......")
	var request AmendPasswordRequest

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&request)

	myuser, ok := userData[request.Username]
	if !ok {
		fmt.Fprintln(w, "该用户不存在!!")
		return
	}

	if myuser.Password != request.OldPassword {
		fmt.Fprintln(w, "密码输入错误!!")
		return
	}

	myuser.Password = request.NewPassword
	userData[myuser.Name] = myuser

	fmt.Fprintln(w, "修改成功!")
}

//获取用户信息函数
func See(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")

	if name == "" {
		fmt.Fprintln(w, "请输入用户名!!")
	}

	user, ok := userData[name]
	if !ok {
		fmt.Fprintln(w, "该用户不存在!!")
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	w.Write(data)
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:1024",
	}

	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/amend/name", AmendName)
	http.HandleFunc("/amend/password", AmendPassword)
	http.HandleFunc("/see", See)
	server.ListenAndServe()

}
