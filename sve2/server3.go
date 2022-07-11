package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloHandler3)
	http.ListenAndServe(":8088", nil)
}

func HelloHandler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello I am 8088")
}
