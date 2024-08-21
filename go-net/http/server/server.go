package main

import (
	"fmt"
	"net/http"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%#v\n", *r.URL)
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>hello,world</h1>")
}

func main() {
	http.HandleFunc("/hello", HandleHello)
	fmt.Println("Listen :8080...")
	http.ListenAndServe(":8080", nil)
}
