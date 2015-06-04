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
	server := &http.Server{Addr: "172.31.21.68:58888", Handler: nil}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Start http server failed: %v", err.Error())
	}
	fmt.Println("Server stopping")
}
