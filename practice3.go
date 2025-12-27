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
			fmt.Println("BIG!")
		} else if input < number {
			fmt.Println("SMALL!!")
		} else if input == number {
			fmt.Println("CLEAR!!")
			os.Exit(0)
		} else {
			fmt.Println("!?!?!?")
		}
	}
	fmt.Println("game over")
	fmt.Println("answer : ", number)
}
