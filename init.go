package main

import r "github.com/lachee/raylib-goplus/raylib"

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
}
