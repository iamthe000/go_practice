package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	suuji := [2]string{"0", "1"}
	kazu := ""

	for x > 0 {
		z := suuji[x%2]
		kazu = z + kazu
		x = x / 2
	}

	fmt.Println(kazu)
}
