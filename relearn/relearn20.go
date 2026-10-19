package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 15; i++ {
		if i%2 != 0 { //「i%2」iを2で割ったあまり　「!=0」0ではない（奇数判定）
			fmt.Println(i)
		} else {
			fmt.Print()
		}
	}
}
