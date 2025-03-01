package main

import (
	"log"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello world!\n"))
}

func main() {
	http.HandleFunc("/hello", HelloWorldHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
