package game

import (
	"fmt"

	"github.com/RaniSputnik/ko/model"
)

var adminUser = &model.User{
	Username: "admin",
}

// Action represents a mutation to the game.
//
// Actions can represent moves from the player,
// chat messages, system events and skips / resignations.
//
// String will return a description of the action.
// TODO consider localization.
// Actor returns the user who caused the action.
// TODO add time of the action
type Action interface {
	fmt.Stringer
	Actor() *model.User
}

// PlayStone is an Action that indicates the playing
// of a stone on the board.
type PlayStone struct {
	player *model.User
	X, Y   int
}

// Actor returns the player who played the stone.
func (mv PlayStone) Actor() *model.User {
	return mv.player
}

func (mv PlayStone) String() string {
	return fmt.Sprintf("%s played a stone at %d,%d", mv.player.Username, mv.X, mv.Y)
}

// Position returns the x,y position of the stone on the board
func (mv PlayStone) Position() (x, y int) {
	return mv.X, mv.Y
}

// Skip is an Action that indicates a player skipped
// their turn.
type Skip struct {
	player *model.User
}

// Actor returns the player who skipped their turn.
func (mv Skip) Actor() *model.User {
	return mv.player
}

func (mv Skip) String() string {
	return fmt.Sprintf("%s skipped their turn.", mv.player.Username)
}
