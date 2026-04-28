package entities

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	id    int16
	pos_x float32
	pos_y float32
	size  float32

	vel_x        float32
	vel_y        float32
	acceleration float32

	max_speed      float32
	angle          float32
	rotation_speed float32

	friction float32

	health int16
}

func CreatePlayer() *Player {
	var player Player = Player{
		id:    0,
		pos_x: 200,
		pos_y: 200,
		size:  20,

		vel_x:        0,
		vel_y:        0,
		acceleration: 1000,

		max_speed:      1000,
		angle:          0,
		rotation_speed: 3.0,

		friction: 0.999,

		health: 3,
	}

	return &player
}

func UpdatePlayer(player *Player) {
	drawPlayer(*player)
	movement(player)
}

func PrintStatus(player Player) {
	fmt.Println(player)
}

func movement(player *Player) {
	rotatePlayer(player)

	goFowardAndBackPlayer(player)

	hitBorderMap(player)
}

func rotatePlayer(player *Player) {
	dt := rl.GetFrameTime()

	if rl.IsKeyDown(rl.KeyA) {
		player.angle -= player.rotation_speed * dt
	}

	if rl.IsKeyDown(rl.KeyD) {
		player.angle += player.rotation_speed * dt
	}
}

func goFowardAndBackPlayer(player *Player) {
	dt := rl.GetFrameTime()
	dirX := float32(math.Cos(float64(player.angle)))
	dirY := float32(math.Sin(float64(player.angle)))

	if rl.IsKeyDown(rl.KeyW) {
		player.vel_x += player.acceleration * dt
		player.vel_y += player.acceleration * dt
	}

	if rl.IsKeyDown(rl.KeyS) {
		player.vel_x -= player.acceleration * dt
		player.vel_y -= player.acceleration * dt
	}

	player.vel_x *= player.friction
	player.vel_y *= player.friction

	speed := float32(math.Sqrt(float64(player.vel_x*player.vel_x + player.vel_y*player.pos_y)))

	if speed > player.max_speed {
		player.vel_x = (player.vel_x / speed) * player.max_speed
		player.vel_y = (player.vel_y / speed) * player.max_speed
	}

	player.pos_x += dirX * player.vel_x * dt
	player.pos_y += dirY * player.vel_y * dt
}

func hitBorderMap(player *Player) {
	if player.pos_x < 0 {
		player.pos_x = float32(rl.GetScreenWidth())
	}

	if player.pos_x > float32(rl.GetScreenWidth()) {
		player.pos_x = 0.0
	}

	if player.pos_y < 0 {
		player.pos_y = float32(rl.GetScreenHeight())
	}

	if player.pos_y > float32(rl.GetScreenHeight()) {
		player.pos_y = 0.0
	}
}

func drawPlayer(player Player) {

	cos := float32(math.Cos(float64(player.angle)))
	sin := float32(math.Sin(float64(player.angle)))

	// Triangle edges coordinates
	top := rl.Vector2{X: player.size, Y: 0}
	left := rl.Vector2{X: -player.size, Y: -player.size}
	right := rl.Vector2{X: -player.size, Y: player.size}

	rotate := func(v rl.Vector2) rl.Vector2 {
		return rl.Vector2{
			X: v.X*cos - v.Y*sin + player.pos_x,
			Y: v.X*sin + v.Y*cos + player.pos_y,
		}
	}

	p1 := rotate(top)
	p2 := rotate(left)
	p3 := rotate(right)

	rl.DrawTriangle(p1, p2, p3, rl.White)
}
