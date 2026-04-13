package main

import (
    "fmt"
    "strconv"
)

func main() {
    n, err := strconv.Atoi("123")

    if err != nil {
        fmt.Println("エラーが発生しました:", err)
    } else {
        fmt.Println("変換成功！数値は:", n)
    }
}
