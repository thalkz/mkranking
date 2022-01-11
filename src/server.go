package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	fmt.Println("Listening on port", httpPort)
	err := http.ListenAndServe(":"+httpPort, nil)
	if err != nil {
		fmt.Println("err")
	}
}
