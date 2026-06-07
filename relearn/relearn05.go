package main

import "fmt"

func main() {
    count := 0
    for i := 1; i <= 100; i++ {
        if i % 3 == 0 { //i%3==0 : 3で割ってあまりが0のとき
            count++
        }
    }
    fmt.Printf("1から100の間に3の倍数は %d 個あります。\n", count)
}
