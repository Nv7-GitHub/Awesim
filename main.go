package main

import (
	r "github.com/lachee/raylib-goplus/raylib"
)

const size = 100
const quality = 3
const directions = 32
const threshold = 0.2

var waterColor = r.NewVector4(0, 1, 1, 1)

func main() {
	screenWidth := 800
	screenHeight := 450

	r.InitWindow(screenWidth, screenHeight, "raylib [shaders] example - custom uniform variable")
	defer r.UnloadAll()

	// Set values
	shader := r.LoadShader("", "shaders/blur.fs") // Load postpro shader
	shader.SetValueFloat32(shader.GetLocation("size"), []float32{float32(size)}, r.UniformFloat)
	shader.SetValueFloat32(shader.GetLocation("quality"), []float32{float32(quality)}, r.UniformFloat)
	shader.SetValueFloat32(shader.GetLocation("directions"), []float32{float32(directions)}, r.UniformFloat)
	shader.SetValueFloat32(shader.GetLocation("threshold"), []float32{float32(threshold)}, r.UniformFloat)
	shader.SetValueFloat32(shader.GetLocation("waterColor"), waterColor.Decompose(), r.UniformVec4)

	target := r.LoadRenderTexture(screenWidth*2, screenHeight*2)

	r.SetTargetFPS(60)

	for !r.WindowShouldClose() {
		r.SetMouseScale(2, 2)

		//Begin drawing like normal
		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)

		// Render onto texture
		r.BeginTextureMode(target)
		r.ClearBackground(r.RayWhite)
		pos := r.GetMousePosition()
		r.DrawCircle((screenWidth*2)/2, (screenHeight*2)/2, 100, r.GopherBlue)
		r.DrawCircle(int(pos.X), int(pos.Y), 100, r.GopherBlue)
		r.EndTextureMode()

		// Render onto screen
		r.BeginShaderMode(shader)
		r.DrawTextureRec(target.Texture, r.NewRectangle(0, 0, float32(target.Texture.Width), float32(-target.Texture.Height)), r.NewVector2(0, 0), r.White)
		r.EndShaderMode()
		r.DrawFPS(10, 10)

		r.EndDrawing()
	}

	r.CloseWindow()
}
