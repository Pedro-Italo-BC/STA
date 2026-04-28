package main

import (
	ent "github.com/Pedro-Italo-BC/STA/game/entities"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var p *ent.Player = ent.CreatePlayer()
	ent.PrintStatus(*p)

	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(800, 800, "Test")

	defer rl.CloseWindow()
	var y float32 = 100.0
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.DrawTriangle(
			rl.Vector2{X: 400, Y: y},
			rl.Vector2{X: 300, Y: 400},
			rl.Vector2{X: 500, Y: 400},
			rl.Red,
		)

		if rl.IsKeyPressed(rl.KeyW) {
			y += 20
		}
		rl.ClearBackground(rl.Black)
		rl.DrawText("Bunda", 300, 300, 20, rl.White)
		rl.EndDrawing()
	}
}
