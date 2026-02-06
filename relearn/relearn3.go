package main

import "fmt"

func main() {
    hello := func(){
        for i:=1;i<10;i++{
            fmt.Println("Hello, world!")
        }
    }
    for j:=1;j<10;j++{hello()}
}
