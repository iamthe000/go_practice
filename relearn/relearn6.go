package main

import "fmt"

func main() {
    var counter int = 0
    hello := func(){
        for i:=0;i<10;i++{
            fmt.Println("Hello, world!")
            counter++
        }
    }
    for j:=0;j<10;j++{hello()}
    fmt.Println("counter :",counter)
    if counter == 100 {
        fmt.Println("SUCCESS!")
    } else {
        fmt.Println("oh...")
    }
}
