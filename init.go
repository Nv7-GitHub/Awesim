package main

import (
	"github.com/jakecoffman/cp"
	r "github.com/lachee/raylib-goplus/raylib"
)

func preGameInit() {
	shader = r.LoadShaderCode(defaultVs, blurFs)
}

func loadGame() {
	shader.SetValueFloat32(shader.GetLocation("size"), []float32{float32(size)}, r.UniformFloat)
	shader.SetValueFloat32(shader.GetLocation("quality"), []float32{float32(quality)}, r.UniformFloat)
	shader.SetValueFloat32(shader.GetLocation("directions"), []float32{float32(directions)}, r.UniformFloat)
	shader.SetValueFloat32(shader.GetLocation("threshold"), []float32{float32(threshold)}, r.UniformFloat)

	textures = make([]r.RenderTexture2D, len(layers))
	for i := range textures {
		textures[i] = r.LoadRenderTexture(width, height)
	}

	space = cp.NewSpace()
	space.SetGravity(cp.Vector{X: 0, Y: gravity})
	space.SetCollisionSlop(0.5)
	for i := 0; i < len(terrain)-1; i++ {
		a := terrain[i]
		b := terrain[i+1]
		shp := space.AddShape(cp.NewSegment(space.StaticBody, a, b, terrainWidth))
		shp.UserData = LayerTerrain
	}
}
