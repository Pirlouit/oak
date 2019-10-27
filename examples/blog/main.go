package main

import (
	"github.com/Pirlouit/oak"
	"github.com/Pirlouit/oak/examples/blog/components"
	"github.com/Pirlouit/oak/pkg/router"
)

func main() {
	// Starts the Oak framework
	oak.Start()

	// Starts our Router
	router.NewRouter()
	router.RegisterRoute("home", components.Home)
	router.RegisterRoute("about", components.About)

	// keeps our app running
	done := make(chan struct{}, 0)
	<-done
}
