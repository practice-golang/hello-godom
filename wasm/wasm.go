package main // import "hello"

import (
	"syscall/js"

	. "github.com/siongui/godom/wasm"
)

var (
	g = js.Global()
	d = Document
)

func say(this js.Value, i []js.Value) interface{} {
	a := d.GetElementById("name")
	b := d.GetElementById("msg")

	b.SetInnerHTML(a.InnerHTML() + " is Not Paocai")

	return nil
}

func registerCallbacks() {
	g.Set("say", js.FuncOf(say))
}

func main() {
	w := make(chan struct{}, 0)

	registerCallbacks()

	println("WASM Go Initialized")
	<-w
}
