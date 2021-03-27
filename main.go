package main

import (
	_ "embed"

	r "github.com/lachee/raylib-goplus/raylib"
)

func main() {
	r.InitWindow(width/2, height/2, "Awesim")
	defer r.UnloadAll()
	r.SetTargetFPS(60)
	r.SetTraceLogLevel(r.LogError | r.LogDebug | r.LogWarning)

	preGameInit()

	// Initialize game
	loadGame()

	for !r.WindowShouldClose() {
		r.SetMouseScale(2, 2)

		renderGame()
	}

	r.CloseWindow()
}
