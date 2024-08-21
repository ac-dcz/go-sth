package main

import (
	"fmt"
	"net/http"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>hello,world</h1>")
}

func main() {
	http.HandleFunc("/hello", HandleHello)
	fmt.Println("Listen :8080...")
	http.ListenAndServeTLS(":8080", "../x509/server.crt", "../x509/server.key", nil)
}
