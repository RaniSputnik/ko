package model

import (
	"fmt"
)

type Event interface {
	Player() Player
	Message() string
}

type hasPlayer struct{ player Player }

func (p hasPlayer) Player() Player { return p.player }

type PlaceStoneEvent struct {
	hasPlayer
	X, Y int
}

func (e PlaceStoneEvent) Message() string {
	// TODO localize
	return fmt.Sprintf("'%s' played a stone at %d,%d", e.player.Username, e.X, e.Y)
}
