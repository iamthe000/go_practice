package main

import "github.com/fatih/color"

func main() {
    color.Red("赤いメッセージ")
    color.Blue("青いメッセージ")
    c := color.New(color.FgCyan).Add(color.Underline)
    c.Println("シアン色でアンダーライン付きの文字")
}
