package main

import (
	"sync"
	"net/http"
	"fmt"
	"log"
)

//记录访问网站的次数
var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/handler", handler1)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL->path:%q\n", r.URL.Path)
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s , %s , %s \n", r.Method, r.Body, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}

	fmt.Fprintf(w, "Host=%q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr=%q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q]=%q\n", k, v)
	}
}
	//
	//GET , {} , HTTP/1.1
	//Header["User-Agent"]=["Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"]
	//Header["Upgrade-Insecure-Requests"]=["1"]
	//Header["Accept"]=["text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"]
	//Header["Accept-Encoding"]=["gzip, deflate, br"]
	//Header["Accept-Language"]=["zh-CN,zh;q=0.9,en;q=0.8"]
	//Header["Connection"]=["keep-alive"]
	//Host="localhost:8080"
	//RemoteAddr="127.0.0.1:58402"

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count %d\n", count)
	mu.Unlock()
}
