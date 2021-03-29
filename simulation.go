package main

import (
	"github.com/jakecoffman/cp"
	r "github.com/lachee/raylib-goplus/raylib"
)

var oldTime = r.GetTime()
var accumulator float64 = 0

func addParticle(pos cp.Vector, layer int) {
	lr := layers[layer]
	body := space.AddBody(cp.NewBody(lr.Mass, cp.MomentForCircle(lr.Mass, 0, lr.Size, cp.Vector{})))
	body.SetPosition(pos)

	shp := space.AddShape(cp.NewCircle(body, lr.Size, cp.Vector{}))
	shp.SetElasticity(0)
	shp.SetFriction(lr.Friction)
	shp.SetCollisionType(cp.CollisionType(lr.Type))
	shp.UserData = lr.Type
}

func addSegment(pos1 cp.Vector, pos2 cp.Vector, layer int) {
	lr := layers[layer]
	shp := space.AddShape(cp.NewSegment(space.StaticBody, pos1, pos2, lr.Size))
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

	space.EachShape(func(s *cp.Shape) {
		if s.UserData.(LayerType) != LayerTerrain {
			shp := s.Class.(*cp.Circle)
			pos := shp.Body().Position()
			if pos.X > width || pos.X < 0 || pos.Y < 0 || pos.Y > height {
				removeShape(s)
			}
		}
	})
}
