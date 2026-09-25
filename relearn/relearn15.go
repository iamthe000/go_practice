package main

import "fmt"

func main(){

_, err := fmt.Println("test")
if err != nil {
fmt.Println("...lol")
}
//これが基本的なGoのエラー処理。エラーの内容がerrとして出てくるのでそれがnil(無)で無かった時(!=)にする処理を下にぶら下げる

if _, err := fmt.Println("test"); err != nil {
    fmt.Println("...lol")
}
//こうするとerrが使われるのがこのスコープだけになるから、ほかのエラー処理と競合を避けれる(うれしいね)

}