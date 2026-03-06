package main

import "fmt"

func main() {
    var color string
    for {
        fmt.Print("color(red or green or blue) : ")
        fmt.Scan(&color)
        if color == "red" {
            fmt.Println("\x1b[31mこれは赤色の文字です\x1b[0m")
        } else if color == "green" {
            fmt.Println("\x1b[32mこれは緑色の文字です\x1b[0m")
        } else if color == "blue" {
            fmt.Println("\x1b[34mこれは青色の文字です\x1b[0m")
        } else if color == "end" {
            fmt.Println("bye")
            break
        } else {
            fmt.Println("??????????????????????????")
        }
    }
}
