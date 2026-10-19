package main

import (
	"fmt"
)

func main() {
	items := []any{2, 5, 8, 9, 30, 45, 55, 102, 183, 209} // せっかくだから使ったことないanyを使ってみる。多分intが適切
	for i := 0; i < len(items); i++ {
		fmt.Println(items[i])
	}
}
