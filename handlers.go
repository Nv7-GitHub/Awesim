package main

import (
	"math"

	"github.com/jakecoffman/cp"
)

func addHandlers() {
	handler := space.NewCollisionHandler(cp.CollisionType(LayerLava), cp.CollisionType(LayerWater))
	handler.PostSolveFunc = lavaWater

	handler = space.NewCollisionHandler(cp.CollisionType(LayerLava), cp.CollisionType(LayerBomb))
	handler.PostSolveFunc = lavaBomb
}

func lavaWater(arb *cp.Arbiter, space *cp.Space, userData interface{}) {
	a, b := arb.Shapes()

	addParticle(a.Body().Position(), int(LayerStone))

	removeShape(a)
	removeShape(b)
}

func lavaBomb(arb *cp.Arbiter, space *cp.Space, userData interface{}) {
	a, b := arb.Shapes()
	pos := a.Body().Position()

	removeShape(a)
	removeShape(b)

	// BOOM
	space.BBQuery(cp.NewBBForExtents(pos, layers[LayerBomb].Extradata["BombRange"], layers[LayerBomb].Extradata["BombRange"]), cp.SHAPE_FILTER_ALL, func(shape *cp.Shape, data interface{}) {
		layer := shape.UserData.(LayerType)
		if layer != LayerTerrain {
			shp := shape.Class.(*cp.Circle)

			// Math that uses 10 variables because I have a small brain
			mag := layers[LayerBomb].Extradata["BombForce"]
			change := pos.Sub(shp.Body().Position())
			angle := math.Atan(change.Y / change.X)
			angle = angle * 180 / math.Pi // RADIANS BAD, DEGREES GOOD
			xChange := math.Cos(angle) * mag
			yChange := math.Sin(angle) * mag

			shp.Body().ApplyImpulseAtLocalPoint(cp.Vector{X: xChange, Y: yChange}, cp.Vector{X: 0, Y: 0})
		}
	}, nil)
}
