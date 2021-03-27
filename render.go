package main

import (
	"github.com/jakecoffman/cp"

	r "github.com/lachee/raylib-goplus/raylib"
)

func renderGame() {
	r.BeginDrawing()
	r.ClearBackground(r.RayWhite)

	for i := range textures {
		r.BeginTextureMode(textures[i])
		r.ClearBackground(r.RayWhite)

		// Render circles instead
		pos := r.GetMousePosition()
		r.DrawCircle(width/2, height/2, 100, r.GopherBlue)
		r.DrawCircle(int(pos.X), int(pos.Y), 100, r.GopherBlue)

		r.EndTextureMode()
	}

	r.BeginShaderMode(shader)
	for i, tex := range textures {
		shader.SetValueFloat32(shader.GetLocation("inpColor"), layers[i].Color.Decompose(), r.UniformVec4)
		r.DrawTextureRec(tex.Texture, r.NewRectangle(0, 0, float32(tex.Texture.Width), float32(-tex.Texture.Height)), r.NewVector2(0, 0), r.White)
	}
	r.EndShaderMode()

	space.EachShape(func(s *cp.Shape) {
		if s.UserData == LayerTerrain {
			shp := s.Class.(*cp.Segment)
			r.DrawLineEx(cp2r2(shp.A()), cp2r2(shp.B()), float32(terrainWidth*2), r.Black)
		}
	})

	r.DrawFPS(10, 10)
	r.EndDrawing()
}
