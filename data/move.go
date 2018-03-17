package data

import (
	"context"
)

type Move struct {
	UserID  string
	MatchID string
	X       int
	Y       int
}

type MoveStore interface {
	SaveMove(ctx context.Context, move Move) (Move, error)
}

const createMoveQuery = `INSERT INTO Moves (MatchID, Player, MoveX, moveY) VALUES (?,?,?,?)`

func (store MysqlStore) SaveMove(ctx context.Context, move Move) (Move, error) {
	var userID, matchID int64
	var err error
	if userID, err = idToInt(move.UserID); err != nil {
		return Move{}, err
	}
	if matchID, err = idToInt(move.MatchID); err != nil {
		return Move{}, err
	}
	_, err = store.DB.ExecContext(ctx, createMoveQuery, matchID, userID, move.X, move.Y)
	return move, err
}
