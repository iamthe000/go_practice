package main

import "fmt"

func main() {
	x := 13

	suuji := [2]string{"0", "1"}
	kazu := ""

	for x > 0 {
		z := suuji[x%2]
		kazu = z + kazu
		x = x / 2
	}

	fmt.Println(kazu)
}
