package main

import (
	"github.com/jakecoffman/cp"
	r "github.com/lachee/raylib-goplus/raylib"
)

// cp2r2 means ChiPmunk To Raylib vector 2
func cp2r2(v cp.Vector) r.Vector2 {
	return r.NewVector2(float32(v.X), float32(v.Y))
}
