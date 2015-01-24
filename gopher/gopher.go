package main

import (
	"github.com/gophergala/GopherKombat/common/game"
	"log"
)

type UserData struct {
}

type Gopher struct {
	Id int
	// Custom user data
	Data UserData
}

func (gopher *Gopher) Init() {
	log.Printf("Init gopher %d\n", gopher.Id)
}

func (gopher *Gopher) Turn(state *game.State) *game.Action {
	log.Printf("Turn: %#v\n", state)
	return &game.Action{Test: "Test action"}
}
