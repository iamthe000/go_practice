package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("入力してください: ")

	// バイト配列（バッファ）を用意
	buf := make([]byte, 1024)

	// os.Stdin から直接バイト列として読み込む（Rawに近い低レベルAPI）
	n, err := os.Stdin.Read(buf)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}

	// 読み込んだバイト数と、生のバイトデータを表示
	fmt.Printf("読み込んだバイト数: %d\n", n)
	fmt.Printf("生のバイトデータ: %v\n", buf[:n])
	fmt.Printf("文字列に変換: %s", string(buf[:n]))
}
