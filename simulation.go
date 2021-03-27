package main

import (
	"github.com/jakecoffman/cp"
	r "github.com/lachee/raylib-goplus/raylib"
)

var oldTime = r.GetTime()
var accumulator float64 = 0

var tool = 0

func addHandlers() {
	handler := space.NewCollisionHandler(cp.CollisionType(LayerWater), cp.CollisionType(LayerLava))
	handler.PostSolveFunc = func(arb *cp.Arbiter, space *cp.Space, userData interface{}) {
		a, b := arb.Shapes()

		addParticle(a.Body().Position(), int(LayerStone))

		space.RemoveBody(b.Body())
		space.RemoveShape(b)
		space.RemoveBody(a.Body())
		space.RemoveShape(a)
	}
}

func addParticle(pos cp.Vector, layer int) {
	lr := layers[layer]
	body := space.AddBody(cp.NewBody(lr.Mass, cp.MomentForCircle(lr.Mass, 0, lr.Radius, cp.Vector{})))
	body.SetPosition(pos)

	shp := space.AddShape(cp.NewCircle(body, lr.Radius, cp.Vector{}))
	shp.SetElasticity(0)
	shp.SetFriction(lr.Friction)
	shp.SetCollisionType(cp.CollisionType(lr.Type))
	shp.UserData = lr.Type
}

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

	// Input
	if r.IsMouseButtonDown(r.MouseLeftButton) {
		pos := r.GetMousePosition()
		for i := 0; i < 180/r.GetFPS(); i++ {
			addParticle(r22cp(pos), tool)
		}
	}
	if r.IsKeyPressed(r.KeyRight) {
		if tool < len(layers)-1 {
			tool++
		}
	} else if r.IsKeyPressed(r.KeyLeft) {
		if tool > 0 {
			tool--
		}
	}
}
