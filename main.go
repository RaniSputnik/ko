package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/RaniSputnik/ko/data"
	"github.com/RaniSputnik/ko/handle"
	"github.com/RaniSputnik/ko/resolve"
	"github.com/RaniSputnik/ko/resolve/schema"
	"github.com/RaniSputnik/ko/svc"
	_ "github.com/go-sql-driver/mysql"
	"github.com/neelance/graphql-go"
)

func main() {
	host := environmentVariableOrDefault("KO_SQL_HOST", "localhost")
	user := environmentVariableOrDefault("KO_SQL_USER", "root")
	pwd := environmentVariableOrDefault("KO_SQL_PWD", "example")

	db := openDB(host, user, pwd)
	// TODO run migrations should not happen on app startup
	// should be an offline action
	//must(migrations.Up(db, migrations.Dir))

	data := createDataloaders(db)
	s := graphql.MustParseSchema(schema.Text, resolve.Root(data))

	http.Handle("/", handle.GraphiQL("Ko", "/graphql"))
	http.Handle("/graphql", handle.GraphQL(s))

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func environmentVariableOrDefault(name, def string) string {
	if val := os.Getenv(name); val != "" {
		return val
	}
	return def
}

func openDB(host, user, pwd string) *sql.DB {
	var db *sql.DB
	var err error

	maxRetries := 3
	delayDuration := time.Second * 1
	for i := 0; i < maxRetries; i++ {
		log.Printf("Connecting to DB...")
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/ko", user, pwd, host))
		if err == nil {
			log.Printf("Connected to DB successfully!")
			return db
		}
		log.Printf("Failed to connect to DB, retrying in %s", delayDuration)
		<-time.After(delayDuration)
	}
	panic(err)
}

func createDataloaders(db *sql.DB) resolve.Data {
	store := data.MysqlStore{DB: db}
	return resolve.Data{
		MatchSvc: svc.MatchSvc{Store: store},
		PlaySvc:  svc.PlaySvc{MoveStore: store},
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
