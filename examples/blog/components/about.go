package components

import (
	"syscall/js"

	"github.com/Pirlouit/oak"
)

type AboutComponent struct{}

var About AboutComponent

func init() {
	oak.RegisterFunction("coolFunc", CoolFunc)
}

func CoolFunc(this js.Value, i []js.Value) {
	println("does stuff")
}

func (a AboutComponent) Render() string {
	return `<div>
				<h2>About Component Actually Works</h2>
				<button onClick="coolFunc();">Cool Func</button>
			</div>`
}
