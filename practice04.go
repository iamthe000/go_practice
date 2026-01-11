package main

import (
	"fyne.io/fyne/v2"           // Fyneの基本データ型（Size, Window, CanvasObjectなど）
	"fyne.io/fyne/v2/app"       // アプリケーション本体の生成用
	"fyne.io/fyne/v2/container" // レイアウトや部品の配置用
	"fyne.io/fyne/v2/widget"    // ボタンやラベルなどの部品用
)

// AppState はアプリケーション内で使い回すデータやUI部品をまとめる構造体
type AppState struct {
	Label  *widget.Label // テキスト表示用ラベル
	Entry  *widget.Entry // 文字入力用フィールド
	Window fyne.Window   // メインウィンドウの参照
}

func main() {
	// 1. アプリケーションのインスタンスを作成
	myApp := app.New()
	// 2. ウィンドウを作成（引数はタイトルバーの文字）
	myWindow := myApp.NewWindow("Advanced Fyne App")

	// 3. 構造体を初期化し、部品の実体を作成
	state := &AppState{
		Label:  widget.NewLabel("名前を入力してください"),
		Entry:  widget.NewEntry(),
		Window: myWindow,
	}
	// 入力フィールドに薄いグレーのヒント文字を表示
	state.Entry.SetPlaceHolder("ここに文字を入力...")

	// 4. 下で定義した makeUI メソッドを呼び出してレイアウトを取得
	content := state.makeUI()

	// 5. ウィンドウに表示する内容（CanvasObject）をセット
	myWindow.SetContent(content)
	// 6. ウィンドウの初期サイズを指定（幅400px, 高さ300px）
	myWindow.Resize(fyne.NewSize(400, 300))
	// 7. ウィンドウを表示し、アプリケーションを起動（終了までブロック）
	myWindow.ShowAndRun()
}

// makeUI はUI部品をレイアウトして一つの「キャンバスオブジェクト」として返す
func (s *AppState) makeUI() fyne.CanvasObject {
	
	// 「実行」ボタンの作成：クリック時の処理（コールバック）を第2引数に渡す
	submitBtn := widget.NewButton("実行", func() {
		currentInput := s.Entry.Text // 現在入力されている文字列を取得

		// 入力内容に応じた条件分岐（ステート更新）
		switch currentInput {
		case "":
			s.Label.SetText("名前が空欄です！")
		case "admin":
			s.Label.SetText("ADMIN MODE")
		default:
			s.Label.SetText("こんにちは、" + currentInput + "さん！")
		}
	})

	// 「クリア」ボタンの作成
	clearBtn := widget.NewButton("クリア", func() {
		s.Entry.SetText("")            // 入力欄を空にする
		s.Label.SetText("リセットされました") // ラベルの文字を変更
	})

	// ボタン2つを横並び（HBox: Horizontal Box）に配置
	controls := container.NewHBox(submitBtn, clearBtn)

	// 全体を縦並び（VBox: Vertical Box）に配置して返す
	// 上から ラベル -> 入力欄 -> ボタン群 の順
	return container.NewVBox(
		s.Label,
		s.Entry,
		controls,
	)
}
