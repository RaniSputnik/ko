package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"
	_ "github.com/mattes/migrate/source/file"

	"github.com/RaniSputnik/ko/handle"
	"github.com/RaniSputnik/ko/resolve"
	"github.com/RaniSputnik/ko/schema"
	"github.com/RaniSputnik/ko/svc"
	"github.com/neelance/graphql-go"
)

func main() {
	host, user, pwd := "localhost", "root", "example"
	db := openDB(host, user, pwd)
	// TODO run migrations should not happen on app startup
	// should be an offline action
	runMigrations(db)

	data := createDataloaders(db)
	s := graphql.MustParseSchema(schema.Text, resolve.Root(data))

	http.Handle("/", handle.GraphiQL("Ko", "/graphql"))
	http.Handle("/graphql", handle.GraphQL(s))

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func openDB(host, user, pwd string) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/ko", user, pwd, host))
	must(err)
	return db
}

func runMigrations(db *sql.DB) {
	// Create migration driver
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	must(err)
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations", "mysql", driver)
	must(err)

	// Run the migrations
	err = m.Up()
	if err != migrate.ErrNoChange {
		panic(err)
	}
}

func createDataloaders(db *sql.DB) resolve.Data {
	return resolve.Data{
		MatchSvc: svc.MatchSvc{DB: db},
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
