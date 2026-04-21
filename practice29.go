package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, err := strconv.Atoi("12345")
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	fmt.Println(a) // 12345
}

/*
strconvは文字列や数字の変換処理をまとめたパッケージ
Atoi(A to i)はASCII (文字列) を Integer (整数) へ変換する関数
*/
/*
a, _ := strconv.Atoi("12345")
fmt.Println(a)
でエラーを無視した処理も可能
*/