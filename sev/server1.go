package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloHandler1)
	http.ListenAndServe(":8089", nil)
}

func HelloHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello I am 8089")
}
