package main

import (
    "encoding/json" // データJSONに変換するために使用
    "fmt"
    "net/http"      // HTTPサーバーやクライアントの機能を提供
    "strconv"      // 文字列を数値（int）に変換するために使用
    "strings"      // 文字列の切り取りや加工をするために使用
)

// Book は「本」の形を定義した構造体
type Book struct {
    // ID, Title, Author はプログラム内で使う名前
    // `json:"..."` は、JSONになった時の名前（小文字にするのが一般的）
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

// books はメモリ上に保存されたデータのスライス
var books = []Book{
    {ID: 1, Title: "望郷", Author: "湊　かなえ"},
    {ID: 2, Title: "テミスの剣", Author: "中山　七里"},
}

func main() {
    // http.HandleFunc は "このURLに来たら、この関数を呼ぶ" と登録
    // "/books" へのアクセスは、booksHandler 関数を使用
    http.HandleFunc("/books", booksHandler)

    // "/books/" へのアクセス（ID指定など）は、bookDetailHandler 関数を使用
    http.HandleFunc("/books/", bookDetailHandler)

    // サーバーを起動する前に、メッセージを表示します
    fmt.Println("Server started at :8080")


    // もし起動に失敗（ポートが使用中など）した場合は、エラーを返す
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}

// booksHandler : 全件取得の際の関数
func booksHandler(w http.ResponseWriter, r *http.Request) {
    // リクエストのメソッドが GET 以外ならエラーを返す
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // これから送るのはjson形式であることを示す
    w.Header().Set("Content-Type", "application/json")

    // json.NewEncoder(w) で、返信用の口(w)に直接JSONを書き込む準備
    // Encode(books) で、スライスをJSONにして送信
    json.NewEncoder(w).Encode(books)
}

// bookDetailHandler は、IDから特定の1件取得を行う
func bookDetailHandler(w http.ResponseWriter, r *http.Request) {
    // URLのパス（例: "/books/1"）から、"/books/" を取り除いて "1" だけを抽出
    idStr := strings.TrimPrefix(r.URL.Path, "/books/")

    // 抽出した文字列 "1" を、数値の1(intの形)に変換します。失敗したらエラーを返す
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid Book ID", http.StatusBadRequest)
        return
    }

    // データのリスト（books）を一つずつ見て、IDが一致する本を検索
    for _, book := range books {
        if book.ID == id {
            // 一致する本が見つかったら、JSONにして返信
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(book)
            return // 送信完了したので関数を終了
        }
    }

    // 最後まで見つからなかった場合は404 Not Foundを返す
    http.NotFound(w, r)
}
