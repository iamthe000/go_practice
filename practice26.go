package main

import "fmt"

func main() {
    num := 42
    message := "こんにちは"
    isTrue := true

    fmt.Printf("numの型は: %T\n", num)         // int
    fmt.Printf("messageの型は: %T\n", message) // string
    fmt.Printf("isTrueの型は: %T\n", isTrue)   // bool
}
