package model

type User struct {
	ID       string
	Username string
}

type PlayerColour string

const (
	ColourUnknown = PlayerColour("")
	ColourBlack   = PlayerColour("BLACK")
	ColourWhite   = PlayerColour("WHITE")
)

type Player struct {
	User
	Colour PlayerColour
}
