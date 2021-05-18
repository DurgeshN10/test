package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("hello world")
	})
	http.ListenAndServe("localhost:9090", nil)
}
