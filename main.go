package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", ResHello)
	http.HandleFunc("/test", ResHelloX)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ResHello(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println(Soma(10, 20, 30))
		fmt.Println("Hit Get /", time.Now())
		w.Write([]byte("<h1>Hello</h1>"))
	}
}

func ResHelloX(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit Get /test", time.Now())
	if r.Header.Get("Authorization") == "1122" {
		w.Write([]byte("Hello tests"))
	} else {
		w.Write([]byte("Unauthorized"))
	}

}

func Soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}
