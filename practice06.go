package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var u User

	fmt.Println("--- ユーザー登録 ---")
	fmt.Print("名前を入力してください: ")
	fmt.Scan(&u.Name)
	fmt.Print("年齢を入力してください: ")
	fmt.Scan(&u.Age)

	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		fmt.Println("JSONの生成に失敗しました:", err)
		return
	}

	filename := "user.json"
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		fmt.Println("ファイルの書き込みに失敗しました:", err)
		return
	}
	fmt.Printf("\nデータを %s に保存しました。\n", filename)

	fmt.Println("--- 保存された内容の表示 ---")
	readData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("ファイルの読み込みに失敗しました:", err)
		return
	}

	fmt.Println(string(readData))
}
