package main

import "fmt"

func main(){
	test := [2]string{"test","example"}
	fmt.Println(test[0])
	fmt.Println("---------------[test]---------------")
	for _, value := range test {
		fmt.Println(value)
	}
	endless := []string{"apple","banana","cherry","pineapple"}
	fmt.Println("---------------[FRUIT]---------------")
	for _, value := range endless {
		fmt.Println(value)
	}
	fmt.Println("-------------------------------------")
}
