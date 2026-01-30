//====================//
//=bufioを使えなかった=//
//====================//
package main

import "fmt"

func main() {
	var name string
	var age int

	// 名前を聞く
	fmt.Print("お名前を入力してください (例: 田中 太郎): ")
	fmt.Scan(&name) // ここで「田中 太郎」と入力したとする

	// 年齢を聞く
	fmt.Print("年齢を入力してください: ")
	fmt.Scan(&age) // 本来はここで止まって入力を待つはずが...

	fmt.Printf("\n--- 結果 ---\n")
	fmt.Printf("名前: %s\n", name)
	fmt.Printf("年齢: %d\n", age)
}
