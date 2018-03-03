package svc

import (
	"context"
	"database/sql"
	"strconv"

	c "github.com/RaniSputnik/ko/context"
)

const (
	BoardSizeTiny   = 9
	BoardSizeSmall  = 13
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
	return Match{ID: intToID(id), BoardSize: boardSize}, nil
}

const getMatchesQuery = `SELECT MatchID, Owner, BoardSize FROM Matches`

func (svc MatchSvc) GetMatches(ctx context.Context) ([]Match, error) {
	results := []Match{}
	rows, err := svc.DB.QueryContext(ctx, getMatchesQuery)
	if err != nil {
		return results, err
	}
	for rows.Next() {
		var (
			matchID   int64
			owner     int64
			boardSize int
		)

		if err := rows.Scan(&matchID, &owner, &boardSize); err != nil {
			return results, err
		}
		results = append(results, Match{
			ID:        intToID(matchID),
			Owner:     intToID(owner),
			BoardSize: boardSize,
		})
	}
	if err := rows.Err(); err != nil {
		return results, err
	}
	return results, nil
}

func intToID(id int64) string {
	return strconv.FormatInt(id, 10)
}
