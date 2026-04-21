package main

func main() {
    // 普通の構造体
    struct User {
        Name string // 「Name」が名前、「string」が型
    }

    // 無名フィールドを使った構造体
    struct Admin {
        User   // フィールド名を書かず、型名（User）だけを書いている
        Level int
    }
}