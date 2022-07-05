package main

import (
	"net/http"
	"page/controllers/accountcontroller"
	"fmt"
)

const portNumber = ":9000"

func main() {

	http.HandleFunc("/account", accountcontroller.Index)
	http.HandleFunc("/account/index", accountcontroller.Index)
	http.HandleFunc("/account/login", accountcontroller.Login)
	http.HandleFunc("/account/welcome", accountcontroller.Welcome)
	http.HandleFunc("/account/register", accountcontroller.Register)
	http.HandleFunc("/account/logout", accountcontroller.Logout)
	http.HandleFunc("/account/portfolio", accountcontroller.Portfolio)
	
	x := fmt.Sprintf
	fmt.Println(x("Staring application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)

}
