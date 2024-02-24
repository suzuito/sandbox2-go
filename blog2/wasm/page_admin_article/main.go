package main

import (
	"syscall/js"
)

func main() {
	js.Global().Set("hoge", js.ValueOf("This value is set in Go."))
	// main関数が終了すると、Instanceがなくなってしまう。
	// 以下はmain関数の終了を防止するためのブロック。
	blocker := make(chan struct{})
	<-blocker
}
