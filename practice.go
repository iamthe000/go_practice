package main

import "fmt"

func main() {
	var usr_input = ""
	var test string = "test"
	fmt.Print("your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Println(name)
	if name == "user" {
		fmt.Println("hi user")
		fmt.Scan(&usr_input)
		fmt.Println("$: ")
		if usr_input == "hello" {
			fmt.Println("hello!")
		}else{
			fmt.Println("...? more...?")
		}
	} else {
		fmt.Println("who...?")
	}
}
