package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("実行されたプログラム名:", os.Args[0])

	if len(os.Args) > 1 {
		fmt.Println("最初の引数 (os.Args[1]):", os.Args[1])
	} else {
		fmt.Println("引数は入力されていません。")
	}

	fmt.Println("--- 全データ一覧 ---")
	for i, v := range os.Args {
		fmt.Printf("Index %d: %s\n", i, v)
	}
}
