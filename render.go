package main

import (
	"github.com/jakecoffman/cp"

	r "github.com/lachee/raylib-goplus/raylib"
)

func renderGame() {
	r.BeginDrawing()
	r.ClearBackground(r.RayWhite)

	r.BeginTextureMode(tex)
	r.ClearBackground(r.Blank)
	space.EachShape(func(s *cp.Shape) {
		if s.UserData.(LayerType) != LayerTerrain {
			shp := s.Class.(*cp.Circle)
			r.DrawCircleV(cp2r2(shp.Body().Position()), float32(shp.Radius()), layers[int(s.UserData.(LayerType))].Color)
		}
	})
	r.EndTextureMode()

	r.BeginShaderMode(shader)
	r.DrawTextureRec(tex.Texture, r.NewRectangle(0, 0, float32(tex.Texture.Width), float32(-tex.Texture.Height)), r.NewVector2(0, 0), r.White)
	r.EndShaderMode()

	space.EachShape(func(s *cp.Shape) {
		if s.UserData == LayerTerrain {
			shp := s.Class.(*cp.Segment)
			r.DrawLineEx(cp2r2(shp.A()), cp2r2(shp.B()), float32(terrainWidth*2), r.Black)
		}
	})

	toolTxt := "Tool: " + layers[tool].Name
	r.DrawText(toolTxt, width/2-(12*len(toolTxt)), 24, 24, r.Black)

	r.DrawFPS(10, 10)
	r.EndDrawing()
}
