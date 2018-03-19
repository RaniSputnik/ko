package game

import (
	"fmt"

	"github.com/RaniSputnik/ko/model"
)

type Move interface {
	fmt.Stringer
	Player() *model.User
}

type PlayStone struct {
	player *model.User
	X, Y   int
}

func (mv PlayStone) Player() *model.User {
	return mv.player
}

func (mv PlayStone) String() string {
	return fmt.Sprintf("%s played a stone at %d,%d", mv.player.Username, mv.X, mv.Y)
}
