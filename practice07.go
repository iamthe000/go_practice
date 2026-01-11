package main

import "fmt"

// typeはものを後から入れることを前提にしているのね
type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	var u User

	fmt.Println("ユーザー情報を入力してください。")

	fmt.Print("ID: ")
	fmt.Scan(&u.ID)

	fmt.Print("名前: ")
	fmt.Scan(&u.Name)

	fmt.Print("メール: ")
	fmt.Scan(&u.Email)

	fmt.Println("--- 入力された内容 ---")
	fmt.Printf("ID: %d, 名前: %s, メール: %s\n", u.ID, u.Name, u.Email)
}
