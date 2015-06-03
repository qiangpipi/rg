package funs

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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
