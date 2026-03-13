package main
import "fmt"

func main() {
    defer fmt.Println("deferが実行") // 実行を予約するだけ

    fmt.Println("通常の処理")
}