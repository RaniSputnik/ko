package game

const (
	BoardSizeTiny   = 9
	BoardSizeSmall  = 13
	BoardSizeNormal = 19
)

type Board struct {
	Size  int
	Moves []Move
}
