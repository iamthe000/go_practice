package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Server processing at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello! GoのWebサーバーへようこそ!")
}
