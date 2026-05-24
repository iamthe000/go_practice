package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
		<h1>Hello, World! これはGoで作ったウェブサーバーです。</h1>
		<p>こんにちは</p>
		`)
	})

	fmt.Println("サーバーを起動しました: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
	}
}
