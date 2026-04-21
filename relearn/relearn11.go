package main

import "fmt"

func main() {
    finger := 3

    switch finger {
    case 1:
        fmt.Println("親指")
    case 2:
        fmt.Println("人差し指")
    case 3:
        fmt.Println("中指")
    case 4:
        fmt.Println("薬指")
    case 5:
        fmt.Println("小指")
    default:
        fmt.Println("不正な番号です")
    }
}
