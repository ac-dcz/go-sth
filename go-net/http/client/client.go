package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	r, err := http.NewRequest(http.MethodGet, "http://:8080/hello", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	query := r.URL.Query()
	query.Add("name", "dcz")
	query.Add("age", "24")
	r.URL.RawQuery = query.Encode()
	fmt.Println("Default Client")
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Printf("Response erro: %v", err)
	} else {
		s, _ := io.ReadAll(resp.Body)
		fmt.Println(string(s))
	}
}
