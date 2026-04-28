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
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		ent.UpdatePlayer(p)

		rl.ClearBackground(rl.Black)
		rl.EndDrawing()
	}
}
