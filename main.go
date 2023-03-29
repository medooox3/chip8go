package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Coolest Chip-8 Emulator")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose(){
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawText("Raylib is Cool", 250, 200, 32, rl.RayWhite)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}