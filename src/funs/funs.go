package funs

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func Output(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here is in Output")
}

func Proxy(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = ""
	r.ParseForm()
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		http.NotFound(w, r)
		return
	}
	for i, v := range resp.Header {
		for _, k := range v {
			w.Header().Add(i, k)
		}
	}
	_, ok := resp.Header["content-length"]
	if !ok && resp.ContentLength > 0 {
		w.Header().Add("Content-Length", fmt.Sprint(resp.ContentLength))
	}
	w.WriteHeader(resp.StatusCode)
	w.Write(d)
}

func Account(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	f, err := os.Open("./data/account.html")
	if err != nil {
		fmt.Println("File open faild" + err.Error())
	}
	defer f.Close()
	d, _ := ioutil.ReadAll(f)
	w.Write(d)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
}
func ChangePasswd(w http.ResponseWriter, r *http.Request) {
}
func Ant(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	f, err := os.Open("./data/ant.html")
	if err != nil {
		fmt.Println("File open faild" + err.Error())
	}
	defer f.Close()
	d, _ := ioutil.ReadAll(f)
	w.Write(d)
}
