package main

import (
	"fmt"
	"net/http"
	"time"
)

// チェック結果を格納
type Result struct {
	URL     string
	Status  string
	Latency time.Duration
}

func main() {
	http.HandleFunc("/", checkHandler)
	fmt.Println("Server starting at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	// 監視対象のサイトリスト
	targets := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://go.dev",
	}

	// チャネルの作成（Result型のデータをやり取りする）
	resultsChan := make(chan Result)

	// 各URLに対してゴルーチンを起動（並行処理）
	for _, url := range targets {
		go func(u string) {
			start := time.Now()
			resp, err := http.Get(u)

			status := "OK"
			if err != nil {
				status = "Error"
			} else {
				defer resp.Body.Close()
			}

			// チャネルに結果を送信
			resultsChan <- Result{
				URL:     u,
				Status:  status,
				Latency: time.Since(start),
			}
		}(url)
	}

	// 結果の受信と表示
	fmt.Fprintln(w, "--- Website Status Check (Parallel) ---")
	for i := 0; i < len(targets); i++ {
		// チャネルから結果を取り出す（全員分終わるまで待機する）
		res := <-resultsChan
		fmt.Fprintf(w, "URL: %-25s | Status: %-5s | Time: %v\n", res.URL, res.Status, res.Latency)
	}
}
