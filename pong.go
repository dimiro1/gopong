// Copyright 2015 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import "strconv"

// Paddle is the player or the computer
type Paddle struct {
	x      int
	y      int
	w      int
	h      int
	vel    float32
	points int
}

// Paddle Up
func (p *Paddle) Up() {
	p.y -= 1
}

// Paddle Down
func (p *Paddle) Down() {
	p.y += 1
}

// The Ball
type Ball struct {
	x     int
	y     int
	w     int
	h     int
	xVel  int
	yVel  int
	speed int
}

// The game object
// This stores the scores, players, ball
type Pong struct {
	player   *Paddle
	computer *Paddle
	ball     *Ball
}

// Load the game
// Create the necessary objects
func (p *Pong) Load(s Screen) {
	p.player = &Paddle{
		x:      10,
		w:      10,
		h:      30,
		vel:    1,
		points: 0,
	}

	p.player.y = s.Height()/2 - p.player.h/2

	p.computer = &Paddle{
		x:      s.Width() - 20,
		w:      10,
		h:      30,
		vel:    1,
		points: 0,
	}

	p.computer.y = s.Height()/2 - p.computer.h/2

	p.ball = &Ball{
		w:     10,
		h:     10,
		xVel:  1,
		yVel:  -1,
		speed: 1,
	}

	p.reset(s)
}

func (p *Pong) reset(s Screen) {
	p.ball.x = s.Width()/2 - p.ball.w/2
	p.ball.y = s.Height()/2 - p.ball.h/2
}

// Update, Check colisions and update the game state
func (p *Pong) Update(s Screen, keys Keys) {
	// Update Ball position
	p.ball.x = p.ball.x + p.ball.speed*p.ball.xVel
	p.ball.y = p.ball.y + p.ball.speed*p.ball.yVel

	// Check Colisions

	// Walls
	if (p.ball.y + p.ball.h) >= s.Height() {
		p.ball.yVel = -1
	}

	if p.ball.y <= 0 {
		p.ball.yVel = 1
	}

	// Computer
	if (p.ball.x+p.ball.w) >= p.computer.x &&
		(p.ball.y+p.ball.h/2) >= p.computer.y &&
		(p.ball.y) <= (p.computer.y+p.computer.h) {

		p.ball.xVel = -1
	}

	// Player
	if p.ball.x <= (p.player.x+p.player.w) &&
		(p.ball.y+p.ball.h/2) >= p.player.y &&
		p.ball.y <= (p.player.y+p.player.h) {

		p.ball.xVel = 1
	}

	// Dumb Computer AI
	// Middle of the screen
	if p.ball.x > s.Width()/2 {
		if p.ball.y < (p.computer.y + p.ball.h) {
			if p.computer.y >= 0 {
				p.computer.Up()
			}
		} else {
			if p.computer.y <= s.Height()-p.computer.h {
				p.computer.Down()
			}
		}
	}

	// Points

	if p.ball.x > p.computer.x {
		p.player.points++

		p.reset(s)
	}

	if p.ball.x <= 0 {
		p.computer.points++

		p.reset(s)
	}

	// Keyboard
	if keys.IsDown(keyUp) {
		if p.player.y >= 0 {
			p.player.Up()
		}
	}

	if keys.IsDown(keyDown) {
		if p.player.y <= s.Height()-p.player.h {
			p.player.Down()
		}
	}
}

// Draw the objects on the screen
func (p *Pong) Draw(screen Screen) {

	black := Color{0, 0, 0}
	white := Color{255, 255, 255}
	green := Color{0, 200, 0}
	red := Color{200, 0, 0}

	// Background
	screen.DrawRect(0, 0, screen.Width(), screen.Height(), black)

	// Player
	screen.DrawRect(p.player.x, p.player.y, p.player.w, p.player.h, white)

	// Computer
	screen.DrawRect(p.computer.x, p.computer.y, p.computer.w, p.computer.h, red)

	// Ball
	screen.DrawRect(p.ball.x, p.ball.y, p.ball.w, p.ball.h, green)

	// Points
	screen.DrawText(strconv.Itoa(p.player.points), screen.Width()/2-50, 20, white)
	screen.DrawText(strconv.Itoa(p.computer.points), screen.Width()/2+50, 20, white)
}
