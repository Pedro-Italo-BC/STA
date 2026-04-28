package entities

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	pos_x float32
	pos_y float32
	size  float32

	speed        float32
	acceleration float32

	max_speed      float32
	angle          float32
	rotation_speed float32

	friction float32

	health int16
}

func CreatePlayer() *Player {
	var player Player = Player{
		pos_x:          200.0,
		pos_y:          200.0,
		health:         3,
		rotation_speed: 2.0,
		speed:          200.0,
		size:           30.0,
		friction:       1.0,
		angle:          0.0,
	}

	return &player
}

func UpdatePlayer(player *Player) {
	drawPlayer(*player)
	movement(player)
}

func PrintStatus(player Player) {
	var message string = fmt.Sprintf(
		"X: %f\nY: %f\nhealth: %d\nangle: %f",
		player.pos_x,
		player.pos_y,
		player.health,
		player.angle,
	)

	fmt.Println(message)
}

func movement(player *Player) {
	dt := rl.GetFrameTime()

	var foward float32 = 0

	if rl.IsKeyDown(rl.KeyW) {
		foward += 1
	}

	if rl.IsKeyDown(rl.KeyS) {
		foward -= 1
	}

	if rl.IsKeyDown(rl.KeyA) {
		player.angle -= player.rotation_speed * dt
	}

	if rl.IsKeyDown(rl.KeyD) {
		player.angle += player.rotation_speed * dt
	}

	dirX := float32(math.Cos(float64(player.angle)))
	dirY := float32(math.Sin(float64(player.angle)))

	player.pos_x += dirX * player.speed * foward * dt
	player.pos_y += dirY * player.speed * foward * dt
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

	fmt.Println(p1, p2, p3)

	rl.DrawTriangle(p1, p2, p3, rl.White)
}
