// Copyright 2015 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

const (
	KEY_UP   = 38
	KEY_DOWN = 40
)

// Keys help to know which keya are pressed
type Keys struct {
	Pressed map[int]bool
}

// NewKeys create a new Key
func NewKeys() Keys {
	keys := Keys{}
	keys.Pressed = make(map[int]bool)

	return keys
}

// IsDown verify if the given key is pressed
func (k *Keys) IsDown(key int) bool {
	stat, exists := k.Pressed[key]

	return exists && stat
}

// Attached to OnKeyDown event
func (k *Keys) OnKeyDown(key int) {
	k.Pressed[key] = true
}

// Attached to OnKeyUp event
func (k *Keys) OnKeyUp(key int) {
	k.Pressed[key] = false
}
