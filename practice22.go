package main

import (
    "fmt"
    "strings"
)

func main() {
    test := strings.Replace("hello", "l", "L", 1)
    fmt.Println(test)
}
