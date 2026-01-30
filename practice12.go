package main

import (
	"fmt"
)

func main() {
	var word1, word2 string

	fmt.Println("入力を待機中 (例: test test):")

	// スペース区切りで2つの単語を読み込む
	_, err := fmt.Scan(&word1, &word2) //&word1, &word2でそれぞれのポインタを指定している。golangのScanは半角スペースをもとから区切りとして使うので、順番に変数に格納されていく。区切りたくないときはbufioを使う
	if err != nil {
		fmt.Println("入力の読み込みに失敗しました:", err)
		return
	}

	// 判定ロジック
	if word1 == "test" && (word2 == "test" || word2 == "example") {
		fmt.Printf("一致しました: %s %s\n", word1, word2)
	} else {
		fmt.Printf("一致しませんでした: %s %s\n", word1, word2)
	}
}
