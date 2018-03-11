package data

import (
	"context"
	"database/sql"

	"github.com/RaniSputnik/ko/model"
)

type Store interface {
	CreateMatch(ctx context.Context, match model.Match) (model.Match, error)
	GetMatches(ctx context.Context, userID string) ([]model.Match, error)
}

type MysqlStore struct {
	DB *sql.DB
}

const createMatchQuery = `INSERT INTO Matches (Owner, BoardSize) VALUES (?,?)`

func (store MysqlStore) CreateMatch(ctx context.Context, match model.Match) (model.Match, error) {
	userID, err := idToInt(match.Owner)
	if err != nil {
		return model.Match{}, err
	}
	rows, err := store.DB.ExecContext(ctx, createMatchQuery, userID, match.BoardSize)
	if err != nil {
		return model.Match{}, err
	}
	id, err := rows.LastInsertId()
	if err != nil {
		return model.Match{}, err
	}
	match.ID = intToID(id)
	return match, nil
}

const getMatchesQuery = `SELECT MatchID, Owner, BoardSize FROM Matches WHERE Owner = ?`

func (store MysqlStore) GetMatches(ctx context.Context, userID string) ([]model.Match, error) {
	results := []model.Match{}
	rows, err := store.DB.QueryContext(ctx, getMatchesQuery, userID)
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
		results = append(results, model.Match{
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
