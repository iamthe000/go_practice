package main

import (
    "github.com/fatih/color"
    "bufio"
    "fmt"
    "os"
    )
func main() {
    fmt.Println("hello")
    fmt.Print("色を選択してください(赤(1),青(2)):")
    scanner := bufio.NewScanner(os.Stdin)
    
    if scanner.Scan(){
        input := scanner.Text()
        fmt.Println("指定された色:",input)
        if input == "1" {
            color.Red("これは赤です。")
        } else if input == "2" {
            color.Blue("これは青です。")
        } else {
            c := color.New(color.FgCyan).Add(color.Underline)
            c.Println("赤と青以外です。")
        }
    }
}
