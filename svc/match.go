package svc

import (
	"context"
	"database/sql"
	"strconv"

	c "github.com/RaniSputnik/ko/context"
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

const createMatchQuery = `INSERT INTO Matches (Owner) VALUES (?)`

func (svc MatchSvc) CreateMatch(ctx context.Context) (Match, error) {
	userID := c.GetUser(ctx)
	rows, err := svc.DB.ExecContext(ctx, createMatchQuery, userID.ID)
	if err != nil {
		return Match{}, err
	}
	id, err := rows.LastInsertId()
	if err != nil {
		return Match{}, err
	}
	return Match{ID: strconv.FormatInt(id, 10)}, nil
}
