package player

import (
	"fmt"
)

type player struct {
	name string
	life int16
	strength int16
}

type enemy struct {
	player *player
}

func NewPlayer(name string, life int16, strength int16) *player {
	if len(name) == 0 && life < 1 {
		return nil
	}
	var newPlayer = new(player)
	newPlayer.name = name
	newPlayer.life = life
	newPlayer.strength = strength
	return newPlayer
}

func NewEnemy(name string, life int16, strength int16) *enemy {
	if len(name) == 0 && life < 1 {
		return nil
	}
	var newPlayer = NewPlayer(name, life, strength)
	var newEnemy = new(enemy)
	newEnemy.player = newPlayer
	return newEnemy
}

func (this *player) SayName() {
	fmt.Println(this.name)
}

func (this enemy) String() string {
	return this.player.name
}

func init() {
	fmt.Println("Loading player package!")
}
