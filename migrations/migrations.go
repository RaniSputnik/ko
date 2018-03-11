package migrations

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"
	_ "github.com/mattes/migrate/source/file"
)

// Dir is the directory where the database migration files reside
// prefixed with the file:// protocol.
const Dir = "file://./migrations"

func driver(db *sql.DB, migrationsPath string) (*migrate.Migrate, error) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(
		migrationsPath, "mysql", driver)
}

func Up(db *sql.DB, migrationsPath string) error {
	if m, err := driver(db, migrationsPath); err != nil {
		return err
	} else if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	} else {
		return nil
	}
}

func Down(db *sql.DB, migrationsPath string) error {
	if m, err := driver(db, migrationsPath); err != nil {
		return err
	} else if err = m.Down(); err != nil && err != migrate.ErrNoChange {
		return err
	} else {
		return nil
	}
}
