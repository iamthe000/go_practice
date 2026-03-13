package main

import (
	"fmt"
	"strings"
)

func main() {
	result := strings.Split("a,b,c", ",")
	fmt.Println(result)
}