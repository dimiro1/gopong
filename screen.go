// Copyright 2015 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

// A color object
type Color struct {
	R int // Red
	G int // Green
	B int // Blue
}

// The Screen interface
// It hides the internal implementation of the screen drawing
type Screen interface {
	// Draw a rectangle on the screen
	DrawRect(x, y, w, h int, color Color)
	// Draw the text at the position
	DrawText(text string, x, y int, color Color)
	// The Screen width
	Width() int
	// The screen Height
	Height() int
}

// A HTML5 Canvas Screen implementation
type CanvasScreen struct {
	Canvas *js.Object // Javascript canvas Object
}

// HTML5 Canvas implementation of a DrawText
func (s *CanvasScreen) DrawText(text string, x, y int, color Color) {
	ctx := s.Canvas.Call("getContext", "2d")

	ctx.Set("fillStyle", fmt.Sprintf("rgb(%d, %d, %d)", color.R, color.G, color.B))
	ctx.Call("fillText", text, x, y)
}

// HTML5 Canvas implementation of a DrawRect
func (s *CanvasScreen) DrawRect(x, y, w, h int, color Color) {
	ctx := s.Canvas.Call("getContext", "2d")

	ctx.Set("fillStyle", fmt.Sprintf("rgb(%d, %d, %d)", color.R, color.G, color.B))
	ctx.Call("fillRect", x, y, w, h)
}

// Canvas Width
func (s *CanvasScreen) Width() int {
	return s.Canvas.Get("width").Int()
}

// Canvas Height
func (s *CanvasScreen) Height() int {
	return s.Canvas.Get("height").Int()
}
