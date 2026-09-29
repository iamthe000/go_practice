package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 変換後の数値（n）と、失敗した場合のエラー（err）を同時に受け取る
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("変換に失敗しました:", err)
		return
	}

	// 成功した場合、数値として使える
	fmt.Println(n + 10) // 133
}
