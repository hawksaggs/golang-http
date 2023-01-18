package main

import (
	"github.com/hawksaggs/golang-http/hello"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(hello.SayHello()))
	})

	http.ListenAndServe(":3000", nil)
}
