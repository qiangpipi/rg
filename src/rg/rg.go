package main

import (
	"fmt"
	"funs"
	"net/http"
)

func main() {
	fmt.Println("Starting server")
	//define request handler
	http.HandleFunc("/", funs.Account)
	http.HandleFunc("/createaccount", funs.CreateAccount)
	http.HandleFunc("/changepasswd", funs.ChangePasswd)
	http.HandleFunc("/myant", funs.Ant)
	//start http server
	server := &http.Server{Addr: "10.10.31.95:8888", Handler: nil}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Start http server failed")
	}
	fmt.Println("Server stopping")
}
