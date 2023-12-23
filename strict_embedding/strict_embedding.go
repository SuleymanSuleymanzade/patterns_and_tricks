package main

import "fmt"

type Unit interface {
	Move(x int, y int)
	Teleport(x int, y int)
}

type Position struct {
	x int
	y int
}

func (self *Position) Move(x int, y int) {
	self.x += x
	self.y += y
}

func (self *Position) Teleport(x int, y int) {
	self.x = x
	self.y = y
}

type Player struct {
	*Position
}

func NewPlayer() *Player {
	return &Player{
		Position: &Position{},
	}
}

type Enemy struct {
	*Position
}

func NewEnemy() *Enemy {
	return &Enemy{
		Position: &Position{},
	}
}

func NewUnit(unitType string) Unit {
	switch unitType {
	case "Player":
		return NewPlayer()
	case "Enemy":
		return NewEnemy()
	default:
		return nil
	}
}

func main() {
	player := NewUnit("Player")
	player.Move(2, 4)

	if playerInstance, ok := player.(*Player); ok {
		fmt.Println(playerInstance.Position)
	} else {
		fmt.Println("Not a Player instance")
	}

}
