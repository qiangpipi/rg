package main

import (
	"fmt"
	"funs"
	"net/http"
)

func main() {
	fmt.Println("Starting server")
	//define request handler
	http.HandleFunc("/new", funs.Output)
	http.HandleFunc("/", funs.Proxy)
	//start http server
	server := &http.Server{Addr: "192.169.100.81:8888", Handler: nil}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Start http server failed")
	}
	fmt.Println("Server stopping")
}
