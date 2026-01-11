package main

import (
	"fmt"
	"math/rand/v2"
	"os"
)

func main() {
	number := rand.IntN(100)
	for i := 1; i <= 10; i++ {
		var input int
		fmt.Print("your prediction : ")
		fmt.Scan(&input)

		if input > number {
			fmt.Println("それより小さいです")
			fmt.Println("rest", 10-i)
		} else if input < number {
			fmt.Println("それより大きいです")
			fmt.Println("rest", 10-i)
		} else if input == number {
			fmt.Println("CLEAR!!")
			os.Exit(0)
		} else {
			fmt.Println("!?!?!?")
			fmt.Println("rest", 10-i)
		}
	}
	fmt.Println("game over")
	fmt.Println("answer : ", number)
}
