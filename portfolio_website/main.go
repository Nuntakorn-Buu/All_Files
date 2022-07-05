package main

//ทำการ import library ที่จำเป็น 
import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
	"github.com/raihaninfo/golangLoginSystem/views"
)

// สร้างตัวแปรที่จะเก็บข้อมูล
var store = sessions.NewCookieStore([]byte("secret-password"))

// สร้างฟังก์ชั่นสำหรับเรียกใช้ database
func init() {
	dbcon()
}

// ส่วนของการเรียกใช้งานหลัก
func main() {
	homeView = views.NewView("views/fron-end/index.gohtml")
	aboutView = views.NewView("views/fron-end/about.gohtml")
	portfolioView = views.NewView("views/fron-end/portfolio.gohtml")
	notFountView = views.NewView("views/fron-end/notfount.gohtml")
	loginView = views.NewView("views/fron-end/login.gohtml")
	signupView = views.NewView("views/fron-end/signin.gohtml")
	forgotPassView = views.NewView("views/fron-end/forgotpass.gohtml")
	fotgotAuthView = views.NewView("views/fron-end/forgetauth.gohtml")
	fotgotAuthErrorView = views.NewView("views/fron-end/forgotAuthError.gohtml")
	updatePassView = views.NewView("views/fron-end/updatepass.gohtml")

	r := mux.NewRouter()
	r.PathPrefix("/asset/").Handler(http.StripPrefix("/asset/", http.FileServer(http.Dir("views/fron-end/asset"))))
	r.NotFoundHandler = http.HandlerFunc(notFount)
	r.HandleFunc("/login", login)
	r.HandleFunc("/loginauth", loginAuth)
	r.HandleFunc("/", home)
	r.HandleFunc("/about", about)
	r.HandleFunc("/portfolio", portfolio)
	r.HandleFunc("/signup", signup)
	r.HandleFunc("/signupauth", signupAuth)
	r.HandleFunc("/logout", logout)
	r.HandleFunc("/forgot_pass", forgotPass)
	r.HandleFunc("/forgot_pass_auth", forgotPassAuth)
	r.HandleFunc("/code_verify", forgotCodeVerify)
	r.HandleFunc("/checkpass", checkPass)

	fmt.Println("Listening port :7000")
	http.ListenAndServe(":7000", r)
}