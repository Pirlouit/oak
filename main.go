package oak

import (
	"syscall/js"

	"github.com/Pirlouit/oak/pkg/http"
	"github.com/Pirlouit/oak/pkg/utils"
)

func registerCallbacks() {
	// packages
	utils.RegisterCallbacks()
	http.RegisterCallbacks()
}

func RegisterFunction(funcName string, myfunc func(this js.Value, i []js.Value) interface{}) {
	js.Global().Set(funcName, js.FuncOf(myfunc))
}

func Start() {
	println("Oak Framework Initialized")
	registerCallbacks()
}
