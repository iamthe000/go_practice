package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// 0 から 99 までの整数
	n := rand.IntN(100)
	fmt.Println(n)

	// 0.0 から 1.0 未満の浮動小数点数
	f := rand.Float64()
	fmt.Println(f)

	// 10 から 50 までの数値
	val := 10 + rand.IntN(41) // 10 + [0-40]
	fmt.Println(val)

	// スライスの中からランダムに
	items := []string{"apple", "banana", "cherry"}
	randomIndex := rand.IntN(len(items))
	fmt.Println(items[randomIndex])

}
