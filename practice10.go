package main

import (
    "fmt"
    "os"
    "os/exec"
    )

func main() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
    var input string

    for {
        fmt.Print("情報を入力してください: ")
        fmt.Scan(&input)

        if input != "" {
            break
        }
        
        fmt.Println("入力が空です。もう一度入力してください。")
    }

    fmt.Println("--------------ok!--------------")
    fmt.Println("入力された内容:", input)
}
