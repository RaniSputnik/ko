package svc

import (
	"context"
	"database/sql"
	"strconv"

	c "github.com/RaniSputnik/ko/context"
)

const (
	BoardSizeNormal = 19
)

type MatchSvc struct {
	DB *sql.DB
}

type Match struct {
	ID        string
	Owner     string
	Opponent  string
	BoardSize int
}

const createMatchQuery = `INSERT INTO Matches (Owner, BoardSize) VALUES (?,?)`

func (svc MatchSvc) CreateMatch(ctx context.Context, boardSize int) (Match, error) {
	userID := c.GetUser(ctx)
	rows, err := svc.DB.ExecContext(ctx, createMatchQuery, userID.ID, boardSize)
	if err != nil {
		return Match{}, err
	}
	id, err := rows.LastInsertId()
	if err != nil {
		return Match{}, err
	}
	return Match{ID: strconv.FormatInt(id, 10), BoardSize: boardSize}, nil
}
