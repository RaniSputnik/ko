package migrations_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/RaniSputnik/ko/data/migrations"
)

const TestDir = "file://."

// MySQL server must be reachable in order for these tests to run.
// Should be running on localhost port 3306.
// See README.md for instructions to run mysql Dockerfile.

func TestMigrationsRunSuccessfully(t *testing.T) {
	// TODO pass environment to tests
	user := "root"
	pwd := "example"
	host := "localhost"
	dbName := fmt.Sprintf("ko_migrationstest")

	err := withDbNamed(dbName, user, pwd, host, func(db *sql.DB) {
		if err := migrations.Up(db, TestDir); err != nil {
			t.Errorf("migrations.Up failed: %s", err)
		}

		// TODO is there anyway to avoid the lock here?
		// this always fails with "migrations.Down failed: can't acquire lock"
		if err := migrations.Down(db, TestDir); err != nil {
			t.Errorf("migrations.Down failed: %s", err)
		}
	})

	if err != nil {
		t.Fatal(err)
	}
}

func withDbNamed(dbName, user, pwd, host string, do func(db *sql.DB)) error {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/", user, pwd, host)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return fmt.Errorf("Could not open DB connection to test migrations: %s", err)
	}
	defer db.Close()

	if _, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName); err != nil {
		return fmt.Errorf("Failed to create DB: %s", err)
	}

	defer func() error {
		if _, err := db.Exec("DROP DATABASE IF EXISTS " + dbName); err != nil {
			return fmt.Errorf("Failed to drop DB: %s", err)
		}
		return nil
	}()

	db2, _ := sql.Open("mysql", connectionString+dbName)
	defer db2.Close()
	do(db2)

	return nil
}
