package main

import r "github.com/lachee/raylib-goplus/raylib"

var oldTime = r.GetTime()
var accumulator float64 = 0

func simulateGame() {
	tickRate := 1 / float64(fps)
	now := r.GetTime()
	dt := now - oldTime
	oldTime = now

	accumulator += dt
	for accumulator > tickRate {
		space.Step(tickRate)
		accumulator -= tickRate
	}
}
