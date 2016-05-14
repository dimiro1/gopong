// Copyright 2015 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	// Get the document
	document := js.Global.Get("document")

	// Create the canvas element
	canvas := document.Call("createElement", "canvas")
	canvas.Set("width", 300)
	canvas.Set("height", 200)

	body := document.Get("body")

	// append the canvas element to the html page body
	body.Call("appendChild", canvas)

	// Create the game
	pong := Pong{}

	keys := NewKeys()

	// A canvas Screen
	screen := &CanvasScreen{Canvas: canvas}

	// Initialize Objects
	pong.Load(screen)

	// The loop variable
	var loop func()

	// Called every time frame
	loop = func() {
		pong.Update(screen, keys)
		pong.Draw(screen)
		js.Global.Call("requestAnimationFrame", loop)
	}

	loop()

	// Attach keyboard events
	body.Call("addEventListener", "keydown", func(e js.Object) {
		keys.OnKeyDown(e.Get("keyCode").Int())
	})

	body.Call("addEventListener", "keyup", func(e js.Object) {
		keys.OnKeyUp(e.Get("keyCode").Int())
	})

}
