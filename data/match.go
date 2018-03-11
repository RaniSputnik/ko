package data

import (
	"context"
	"database/sql"
	"errors"

	"github.com/RaniSputnik/ko/model"
)

var (
	ErrNotFound = errors.New("Not found")
)

type Store interface {
	SaveMatch(ctx context.Context, match model.Match) (model.Match, error)
	GetMatches(ctx context.Context, userID string) ([]model.Match, error)
	GetMatch(ctx context.Context, matchID string) (model.Match, error)
}

type MysqlStore struct {
	DB *sql.DB
}

const createMatchQuery = `INSERT INTO Matches (Owner, BoardSize) VALUES (?,?)`

func (store MysqlStore) createMatch(ctx context.Context, match model.Match) (model.Match, error) {
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

const updateMatchQuery = `UPDATE Matches SET BoardSize = ?, Opponent = ? WHERE MatchId = ?`

func (store MysqlStore) SaveMatch(ctx context.Context, match model.Match) (model.Match, error) {
	if match.ID == "" {
		return store.createMatch(ctx, match)
	}
	if rows, err := store.DB.ExecContext(ctx, updateMatchQuery, match.BoardSize, match.Opponent, match.ID); err != nil {
		return match, err
	} else if nrows, _ := rows.RowsAffected(); nrows == 0 {
		return match, ErrNotFound
	}
	return match, nil
}

const getMatchesQuery = `SELECT MatchID, Owner, Opponent, BoardSize FROM Matches WHERE Owner = ?`

func (store MysqlStore) GetMatches(ctx context.Context, userID string) ([]model.Match, error) {
	results := []model.Match{}
	rows, err := store.DB.QueryContext(ctx, getMatchesQuery, userID)
	if err != nil {
		return results, err
	}
	for rows.Next() {
		match, err := scanMatch(rows)
		if err != nil {
			return results, err
		}
		results = append(results, match)
	}
	if err := rows.Err(); err != nil {
		return results, err
	}
	return results, nil
}

const getMatchQuery = `SELECT MatchID, Owner, Opponent, BoardSize FROM Matches WHERE MatchID = ?`

func (store MysqlStore) GetMatch(ctx context.Context, matchID string) (model.Match, error) {
	match := model.Match{}

	intID, err := idToInt(matchID)
	if err != nil {
		return match, err
	}

	row := store.DB.QueryRowContext(ctx, getMatchQuery, intID)
	return scanMatch(row)
}

type scanner interface {
	Scan(dest ...interface{}) error
}

func scanMatch(row scanner) (match model.Match, err error) {
	var (
		gotID     int64
		owner     int64
		opponent  *int64
		boardSize int
	)

	if err = row.Scan(&gotID, &owner, &opponent, &boardSize); err != nil {
		if err == sql.ErrNoRows {
			return match, ErrNotFound
		}
		return match, err
	}

	match = model.Match{
		ID:        intToID(gotID),
		Owner:     intToID(owner),
		BoardSize: boardSize,
	}
	if opponent != nil {
		match.Opponent = intToID(*opponent)
	}
	return match, err
}
