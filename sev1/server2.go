package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloHandler2)
	http.ListenAndServe(":8087", nil)
}

func HelloHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello I am 8087")
}
