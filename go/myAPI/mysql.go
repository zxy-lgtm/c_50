package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name     string `json : "name"`
	Password string `json : "password"`
	Phone    string `json : "phone"`
}

type ScanUser struct {
	ID       int
	Name     string
	Password string
	Phone    string
}

var db *sql.DB

//连接数据库
func initDB() (err error) {
	var password string
	fmt.Println("password:")
	fmt.Scanln(&password)
	dsn := "root:" + password + "@tcp(127.0.0.1:3306)/users"
	//连接数据库,但并没有检查它的正确性
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	//连接数据库并且检查它的正确性
	err = db.Ping()
	if err != nil {
		return
	}
	return err
}

//在数据库中查找用户名字方法,并且返回ID值作为主键!!,返回密码作为登录操作需求
func (user *User) Scanusername() (int, string, bool) {
	sqlStr := "select *from users where name = ?"

	var nuser ScanUser

	rowres := db.QueryRow(sqlStr, user.Name)

	rowres.Scan(&nuser.ID, &nuser.Name, &nuser.Password, &user.Phone)

	if nuser.Name == "" {
		return 0, "", false
	}

	return nuser.ID, nuser.Password, true
}

//插入用户信息的一个方法
func (user *User) AddUser() error {
	sqlStr := "insert into users(name,password,phone) value(?,?,?)"

	inStmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常:", err)
		return err
	}

	_, err = inStmt.Exec(user.Name, user.Password, user.Phone)
	if err != nil {
		fmt.Println("执行出现异常:", err)
		return err
	}
	return nil
}

//更新用户名字信息的一个方法
func (user *User) UpdateName(newname string) {
	sqlStr := "update users set name = ? where id = ?"
	id, password, ok := user.Scanusername()
	if !ok {
		fmt.Println("该用户不存在!!")
		return
	}

	if password != user.Password {
		fmt.Println("密码错误!!")
		return
	}

	db.Exec(sqlStr, newname, id)
	fmt.Println("Yes!!")
}

//更改用户密码的方法
func (user *User) UpdatePassword(newpassword string) {
	sqlStr := "update users set password = ? where id = ?"
	id, password, ok := user.Scanusername()
	if !ok {
		fmt.Println("该用户不存在!!")
		return
	}

	if password != user.Password {
		fmt.Println("密码错误!!")
		return
	}

	db.Exec(sqlStr, newpassword, id)
	fmt.Println("Yes!!")
}

//删除某位用户的方法
func (user *User) DeleteUser() {
	sqlStr := "delete from users where name = ?"
	_, password, ok := user.Scanusername()
	if !ok {
		fmt.Println("该用户不存在!!")
		return
	}

	if password != user.Password {
		fmt.Println("密码错误!!")
		return
	}
	_, err := db.Exec(sqlStr, user.Name)
	if err != nil {
		panic(err)
	}
	fmt.Println("YES!!")
}

//注册函数
func Signin(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	fmt.Fprintln(w, "你好!新用户!\n注册中......")

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(user)

	_, _, ok := user.Scanusername()
	if !ok {
		fmt.Fprintln(w, "注册成功")
		user.AddUser()
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

	user := &User{}
	user.Name = new.Name

	_, password, ok := user.Scanusername()
	if !ok {
		fmt.Fprintln(w, "该用户不存在!!")
		return
	}

	if new.Password != password {
		fmt.Fprintln(w, "密码输入错误!!")
		return
	}

	fmt.Fprintln(w, "登陆成功!")

}

//修改名字结构体
type AmendNameRequest struct {
	OldName  string `json : "oldname"`
	NewName  string `json : "newname"`
	Password string `json : "password"`
}

//修改名字函数
func AmendName(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "你好!用户!\n修改中......")
	var request AmendNameRequest

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&request)

	user := &User{}
	user.Name = request.OldName

	_, password, ok := user.Scanusername()

	if !ok {
		fmt.Fprintln(w, "该用户不存在!!")
		return
	}

	if password != request.Password {
		fmt.Fprintln(w, "密码错误!!")
		return
	}

	user.Password = password
	user.UpdateName(request.NewName)

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

	user := &User{}
	user.Name = request.Username

	_, password, ok := user.Scanusername()
	if !ok {
		fmt.Fprintln(w, "该用户不存在!!")
		return
	}

	if password != request.OldPassword {
		fmt.Fprintln(w, "密码输入错误!!")
		return
	}

	user.Password = password
	user.UpdatePassword(request.NewPassword)

	fmt.Fprintln(w, "修改成功!")
}

//获取用户信息函数
func See(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")

	if name == "" {
		fmt.Fprintln(w, "请输入用户名!!")
	}

	user := &User{}
	user.Name = name
	_, password, ok := user.Scanusername()
	if !ok {
		fmt.Fprintln(w, "该用户不存在!!")
		return
	}
	user.Password = password

	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	w.Write(data)
}

//删除用户结构体
type DeleteRequest struct {
	Name     string `json:"name"`
	Password string `json: "password"`
}

//删除(注销)用户的函数
func Delete(w http.ResponseWriter, r *http.Request) {
	new := new(DeleteRequest)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(new)

	user := &User{}
	user.Name = new.Name

	_, password, ok := user.Scanusername()
	if !ok {
		fmt.Fprintln(w, "该用户不存在!!")
		return
	}

	if new.Password != password {
		fmt.Fprintln(w, "密码输入错误!!")
		return
	}

	user.Password = password
	user.DeleteUser()
	fmt.Fprintln(w, "注销成功!")

}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:2048",
	}

	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/amend/name", AmendName)
	http.HandleFunc("/amend/password", AmendPassword)
	http.HandleFunc("/see", See)
	http.HandleFunc("/delete", Delete)

	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err : %v\n", err)
		return
	}
	fmt.Println("连接数据库成功!")

	server.ListenAndServe()

}
