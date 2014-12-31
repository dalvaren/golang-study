package player

import (
	"fmt"
)

type player struct {
	name string
	life int16
	strength int16
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

func (this *player) SayName() {
	fmt.Println(this.name)
}

func init() {
	fmt.Println("Loading player package!")
}
