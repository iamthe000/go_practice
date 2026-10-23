package main

import "fmt"

func main() {
	raitenTime := []int{0, 2, 9, 10}
	kaishiTime := make([]int, 5)

	end := 0

	for i := 0; i <= 3; i++ {
		var kaishi int
		if raitenTime[i] > end {
			kaishi = raitenTime[i]
		} else {
			kaishi = end
		}
		end = kaishi + 4
		kaishiTime[i] = kaishi
	}

	// 各客の開始時刻の出力
	for i := 1; i <= 4; i++ {
		fmt.Printf("%d人目の開始時刻: %d分\n", i, kaishiTime[i])
	}
}
