package main

import r "github.com/lachee/raylib-goplus/raylib"

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
