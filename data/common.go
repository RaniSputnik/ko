package data

import (
	"database/sql"
	"strconv"
)

type Store interface {
	MatchStore
	MoveStore
}

type MysqlStore struct {
	DB *sql.DB
}

type errInvalidID string

func (v errInvalidID) Error() string {
	return string(v)
}

func intToID(id int64) string {
	return strconv.FormatInt(id, 10)
}

func idToInt(id string) (int64, error) {
	return strconv.ParseInt(id, 10, 0)
}
