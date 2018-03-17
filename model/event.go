package model

type Event interface {
	Player() Player
	Message() string
}

type PlaceStoneEvent struct {
	PlayerID string
	X, Y     int
}
