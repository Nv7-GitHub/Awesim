package main

import (
	"strconv"

	"github.com/jakecoffman/cp"

	r "github.com/lachee/raylib-goplus/raylib"
)

func renderGame() {
	r.BeginDrawing()
	r.ClearBackground(r.RayWhite)

	r.BeginTextureMode(tex)
	r.ClearBackground(r.Blank)
	space.EachShape(func(s *cp.Shape) {
		if layers[int(s.UserData.(LayerType))].RenderType != RenderSegment {
			shp := s.Class.(*cp.Circle)
			r.DrawCircleV(cp2r2(shp.Body().Position()), float32(shp.Radius()), layers[int(s.UserData.(LayerType))].Color)
		}
	})
	r.EndTextureMode()

	r.BeginShaderMode(shader)
	r.DrawTextureRec(tex.Texture, r.NewRectangle(0, 0, float32(tex.Texture.Width), float32(-tex.Texture.Height)), r.NewVector2(0, 0), r.White)
	r.EndShaderMode()

	shapeCount := 0
	space.EachShape(func(s *cp.Shape) {
		if layers[int(s.UserData.(LayerType))].RenderType == RenderSegment {
			shp := s.Class.(*cp.Segment)
			r.DrawLineEx(cp2r2(shp.A()), cp2r2(shp.B()), float32(layers[int(s.UserData.(LayerType))].Size*2), r.Black)
		}
		shapeCount++
	})

	countTxt := "Object Count: " + strconv.Itoa(shapeCount)
	r.DrawText(countTxt, width-((fontSize/2+1)*len(countTxt)), fontSize, fontSize, r.Black)

	handleTools()

	r.DrawFPS(10, height-30)
	r.EndDrawing()
}
