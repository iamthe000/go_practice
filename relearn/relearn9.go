package main

import (
    "fmt"
    "os"
)

func main() {
    // ファイルを開く（ファイルポインタとエラーの2つが返される）
    file, err := os.Open("sonnzaishinaiyo.nothing")

    // エラーのチェック
    if err != nil {
        // エラーが発生した場合の処理
        fmt.Printf("エラーが発生しました: %v\n", err)
        return
    }

    // 正常系の処理
    defer file.Close()
    fmt.Println("ファイルを正常に開きました")
}
