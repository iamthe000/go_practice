package main

import "fmt"

func main() {
    // \x1b[31m が赤色、\x1b[0m がリセット
    fmt.Println("\x1b[31mこれは赤色の文字です\x1b[0m")
    fmt.Println("\x1b[32mこれは緑色の文字です\x1b[0m")
    fmt.Println("\x1b[34mこれは青色の文字です\x1b[0m")
}
