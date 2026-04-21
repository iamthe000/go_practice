package main

import "fmt"

func main() {
	x, y := 1, 2
	x, y = y, x // 値の入れ替え

	fmt.Println(x, y)
}