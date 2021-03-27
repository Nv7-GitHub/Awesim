package main

import (
	"github.com/jakecoffman/cp"
	r "github.com/lachee/raylib-goplus/raylib"
)

func loadGame() {
	shaders = make([]r.Shader, len(layers))
	for i, lr := range layers {
		shaders[i] = r.LoadShaderCode(defaultVs, blurFs)
		shaders[i].SetValueFloat32(shaders[i].GetLocation("size"), []float32{float32(size)}, r.UniformFloat)
		shaders[i].SetValueFloat32(shaders[i].GetLocation("quality"), []float32{float32(quality)}, r.UniformFloat)
		shaders[i].SetValueFloat32(shaders[i].GetLocation("directions"), []float32{float32(directions)}, r.UniformFloat)
		shaders[i].SetValueFloat32(shaders[i].GetLocation("threshold"), []float32{float32(threshold)}, r.UniformFloat)
		shaders[i].SetValueFloat32(shaders[i].GetLocation("inpColor"), lr.Color.Decompose(), r.UniformVec4)
	}

	textures = make([]r.RenderTexture2D, len(layers))
	for i := range textures {
		textures[i] = r.LoadRenderTexture(width, height)
	}

	space = cp.NewSpace()
	space.Iterations = uint(iterations)
	space.SetGravity(cp.Vector{X: 0, Y: gravity})
	space.SetCollisionSlop(0.5)
	addHandlers()
	for i := 0; i < len(terrain)-1; i++ {
		a := terrain[i]
		b := terrain[i+1]
		shp := space.AddShape(cp.NewSegment(space.StaticBody, a, b, terrainWidth))
		shp.UserData = LayerTerrain
		shp.SetCollisionType(cp.CollisionType(LayerTerrain))
	}
}
