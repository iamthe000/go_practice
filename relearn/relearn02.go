package main

import "fmt"

func main() {
	fmt.Print("this is ")
	fmt.Println("test")
	var test string = "test"
	fmt.Println("this is", test)

	var usr_age int
	fmt.Print("your age : ")
	fmt.Scan(&usr_age)
	fmt.Println("your age is", usr_age)
}
