package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloWorldHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
