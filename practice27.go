package main

import (
	"fmt"
	"math"
)

func main() {
	x := 16
	result := math.Sqrt(float64(x))//float64という小数点の使える数字の形式に変換
	fmt.Println(result)
}
