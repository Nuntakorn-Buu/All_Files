package accountcontroller

import (
	"html/template"
	"net/http"
	"fmt"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mysession"))

func Login(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	username := request.Form.Get("username")
	email := request.Form.Get("email")
	password := request.Form.Get("password")

	fmt.Println("username: ", username)
	fmt.Println("email: ", email)
	fmt.Println("password: ", password)

	if email == "62050553@go.buu.ac.th" && password == "123" {
		session, _ := store.Get(request, "mysession")
		session.Values["email"] = email
		session.Save(request, response)
		http.Redirect(response, request, "/account/welcome", http.StatusSeeOther)
	
	} else {
		data := map[string]interface{}{
			"err": "Invalid !!!",
		}
		tmp, _ := template.ParseFiles("views/accountcontroller/index.html")
		tmp.Execute(response, data)
	}
}

func Index(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("views/accountcontroller/index.html")
	tmp.Execute(response, nil)
}

func Logout(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	session.Options.MaxAge = -1
	session.Save(request, response)
	http.Redirect(response, request, "/account/index",http.StatusSeeOther)
}

func Register(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("views/accountcontroller/register.html")
	tmp.Execute(response, nil)
}

func Welcome(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	username := session.Values["username"]
	fmt.Println("username: ", username)
	data := map[string]interface{}{
		"username": username,
	}

	tmp, _ := template.ParseFiles("views/accountcontroller/welcome.html")
	tmp.Execute(response, data)
}

func Portfolio(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("views/accountcontroller/portfolio.html")
	tmp.Execute(response, nil)
}