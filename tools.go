package main

import (
	"github.com/jakecoffman/cp"
	r "github.com/lachee/raylib-goplus/raylib"
)

var tool = 0

func handleTools() {
	toolTxt := "Tool: " + tools[tool].Name
	r.DrawText(toolTxt, fontSize, fontSize, fontSize, r.Black)

	if r.IsKeyPressed(r.KeyRight) {
		if tool < len(tools)-1 {
			tool++
		}
	} else if r.IsKeyPressed(r.KeyLeft) {
		if tool > 0 {
			tool--
		}
	}
	if r.IsKeyPressed(r.KeyR) {
		r.UnloadAll()
		loadGame()
	}

	tools[tool].Tool(tools[tool])
}

func uselessTool(Tool) {
	txt := "USELESS TOOL"
	r.DrawText(txt, width/2-((fontSize/2)*len(txt)), fontSize, fontSize, r.Black)
}

func placeTool(tool Tool) {
	toolTxt := "Layer: " + layers[tool.IntData["tool"]].Name
	r.DrawText(toolTxt, width/2-((fontSize/2)*len(toolTxt)), fontSize, fontSize, r.Black)

	if r.IsMouseButtonDown(r.MouseLeftButton) && layers[tool.IntData["tool"]].RenderType == RenderParticle {
		pos := r.GetMousePosition()
		for i := 0; i < particlePlaceSpeed/r.GetFPS(); i++ {
			addParticle(r22cp(pos), tool.IntData["tool"])
		}
	}

	if layers[tool.IntData["tool"]].RenderType == RenderSegment {
		pos := r.GetMousePosition()
		if r.IsMouseButtonPressed(r.MouseLeftButton) {
			if tool.BoolData["hasPlaced"] {
				addSegment(cp.Vector{X: tool.FloatData["p1x"], Y: tool.FloatData["p1y"]}, r22cp(pos), tool.IntData["tool"])
				tool.BoolData["hasPlaced"] = false
			} else {
				tool.BoolData["hasPlaced"] = true
				tool.FloatData["p1x"] = float64(pos.X)
				tool.FloatData["p1y"] = float64(pos.Y)
			}
		} else if tool.BoolData["hasPlaced"] {
			r.DrawLineEx(r.NewVector2(float32(tool.FloatData["p1x"]), float32(tool.FloatData["p1y"])), pos, float32(layers[int(tool.IntData["tool"])].Size*2), r.Black)
		}
	}

	if r.IsKeyPressed(r.KeyUp) {
		if tool.IntData["tool"] < len(layers)-1 {
			tool.IntData["tool"]++
		}
	} else if r.IsKeyPressed(r.KeyDown) {
		if tool.IntData["tool"] > 0 {
			tool.IntData["tool"]--
		}
	}
}
