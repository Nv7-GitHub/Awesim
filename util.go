package main

import (
	"github.com/jakecoffman/cp"
	r "github.com/lachee/raylib-goplus/raylib"
)

// cp2r2 means ChiPmunk To Raylib vector 2
func cp2r2(v cp.Vector) r.Vector2 {
	return r.NewVector2(float32(v.X), float32(v.Y))
}

// r22cp means Raylib vector 2 To ChiPmunk
func r22cp(v r.Vector2) cp.Vector {
	return cp.Vector{X: float64(v.X), Y: float64(v.Y)}
}

func removeShape(shp *cp.Shape) {
	if space.ContainsBody(shp.Body()) {
		space.RemoveBody(shp.Body())
	}
	if space.ContainsShape(shp) {
		space.RemoveShape(shp)
	}
}
