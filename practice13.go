package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 標準入力（キーボード）を一行丸ごと読み込むスキャナーを作る
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("一行入力してください:")
	if scanner.Scan() {
		// ここではスペースがあっても一行丸ごと a に入る
		a := scanner.Text() 
		fmt.Printf("変数aの中身: [%s]\n", a)
	}
}
