package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		fmt.Println("hello world")
	})
	http.ListenAndServe("localhost:9090", nil)
}
