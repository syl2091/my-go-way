package main

import (
	"fmt"
	"log"
	"net/http"
)

// http基础
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func counter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
