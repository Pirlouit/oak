package main

import (
	"syscall/js"

	"github.com/Pirlouit/oak"
)

func mycoolfunc(this js.Value, i []js.Value) {
	println("My Awesome Function")
}

func main() {
	// Starts the Oak framework
	oak.Start()

	// registers custom functions
	oak.RegisterFunction("coolfunc", mycoolfunc)

	// keeps our app running
	done := make(chan struct{}, 0)
	<-done
}
