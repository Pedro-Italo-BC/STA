package entities

import "fmt"

type Player struct {
	pos_x int;
	pos_y int;
	health int;
	velocity float64;
	angle float64;
}

func CreatePlayer() *Player {
	var player Player := {
		pos_x: 0,
		pos_y: 0,
		health: 3,
		velocity: 1.0,
		angle: 0.0
	}

	return &player;
}

func PrintStatus(player Player) {
	var message string = fmt.Stringf("
	X: %d\n
	Y: %d\n
	health: %d\n
	velocity: %f\n
	angle: %f
	", player.pos_x, player.pos_y, player.health, player.velocity, player.angle);

	fmt.Println(message);
}
