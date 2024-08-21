package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	r, err := http.NewRequest(http.MethodGet, "https://:8080/hello", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	rootCrt, err := os.ReadFile("../x509/ca.crt")
	if err != nil {
		fmt.Println(err)
		return
	}
	rootCAS := x509.NewCertPool()
	rootCAS.AppendCertsFromPEM(rootCrt)
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				ServerName: "www.h.dcz.com",
				RootCAs:    rootCAS,
			},
		},
	}

	fmt.Println("Https Client")
	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf("Response erro: %v", err)
	} else {
		s, _ := io.ReadAll(resp.Body)
		fmt.Println(string(s))
	}

	fmt.Println("Default Client")
	resp, err = http.DefaultClient.Do(r)
	if err != nil {
		fmt.Printf("Response erro: %v", err)
	} else {
		s, _ := io.ReadAll(resp.Body)
		fmt.Println(string(s))
	}
}
