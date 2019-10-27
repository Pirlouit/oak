package http

import (
	"syscall/js"
)

func RegisterCallbacks() {
	js.Global().Set("get", js.FuncOf(GetRequest))
}

func GetRequest(this js.Value, i []js.Value) interface{} {
	println("Does nothing")
	return nil
}
