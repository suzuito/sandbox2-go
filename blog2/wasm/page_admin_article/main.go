package main

import (
	"context"
	"fmt"
	"syscall/js"

	"github.com/suzuito/sandbox2-go/blog2/internal/markdown2html"
)

func main() {
	m2h := markdown2html.Markdown2HTMLImpl{}
	js.Global().Set("hoge", js.ValueOf("This value is set in Go."))
	js.Global().Set("generateMarkdown2HTML", js.FuncOf(func(this js.Value, args []js.Value) any {
		md := args[0].String()
		html := ""
		if err := m2h.Generate(context.Background(), md, &html); err != nil {
			// TODO error handling
			fmt.Println(err)
			return nil
		}
		return js.ValueOf(html)
	}))
	js.Global().
		Set("putMarkdown", js.FuncOf(func(this js.Value, args []js.Value) any {
			return nil
		}))
		/*
			js.Global().
				Get("document").
				Call("getElementById", "preview-button").
				Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) any {
					// プレビューボタンが押されたときの処理
					event := args[0]
					isPreviewChecked := event.Get("target").Get("checked").Bool()
					if isPreviewChecked {
						js.Global().
							Get("document").
							Call("getElementById", "markdown-editor").
							Get("style").
							Set("display", "none")
						js.Global().
							Get("document").
							Call("getElementById", "markdown-viewer").
							Get("style").
							Set("display", "block")
					} else {
						js.Global().
							Get("document").
							Call("getElementById", "markdown-editor").
							Get("style").
							Set("display", "block")
						js.Global().
							Get("document").
							Call("getElementById", "markdown-viewer").
							Get("style").
							Set("display", "none")
					}
					return nil
				}))
					js.Global().
						Get("document").
						Call("getElementById", "markdown-editor").
						Call("addEventListener", "input", js.FuncOf(func(this js.Value, args []js.Value) any {
							// マークダウンエディタの入力が変更されたときの処理
							// fmt.Println("input", this)
							// fmt.Println(this.Get("value"))
							html := ""
							if err := m2h.Generate(context.Background(), this.Get("value").String(), &html); err != nil {
								// TODO error handling
								fmt.Println(err)
								return nil
							}
							js.Global().
								Get("document").
								Call("getElementById", "markdown-viewer").
								Set("innerHTML", html)
							return nil
						}))
		*/
		/*
			js.Global().
				Get("document").
				Call("getElementById", "markdown-editor").
				Call("addEventListener", "input", js.FuncOf(func(this js.Value, args []js.Value) any {
					// マークダウンエディタの入力が変更されたときの処理
					fmt.Println("input", this)
					fmt.Println(this.Get("value"))
					html := ""
					if err := m2h.Generate(context.Background(), this.Get("value").String(), &html); err != nil {
						// TODO error handling
						fmt.Println(err)
						return nil
					}
					js.Global().
						Get("document").
						Call("getElementById", "markdown-viewer").
						Set("innerHTML", html)
					return nil
				}))
		*/
	// main関数が終了すると、Instanceがなくなってしまう。
	// 以下はmain関数の終了を防止するためのブロック。
	blocker := make(chan struct{})
	<-blocker
}