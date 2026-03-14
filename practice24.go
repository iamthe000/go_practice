package main

import (
    "fmt"
    "strings"
)

func main() {
    var adminMode bool = true
    if adminMode {
        defer fmt.Println("last")
        test := strings.Replace("hello", "l", "L", 1)
        fmt.Println(test)
    }
    test := strings.Replace("hello", "l", "L", 1)
    fmt.Println(test)
}
