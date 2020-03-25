package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"server/libs/configuration"
	"server/src/sqlite"
)

type FN func(t *testing.T, db *sqlite.Database)

func Seed(ts *testing.T, name string, fn FN) {
	databaseURL := configuration.Load().DatabaseURL
	if databaseURL == "" {
		databaseURL = ":memory:"
	}
	db := sqlite.New(sqlite.Config{URL: databaseURL})
	executeSQL(db, "taghub.sql")
	executeSQL(db, "_repos.sql")
	ts.Run(name, func(t *testing.T) {
		fn(t, db)
	})
}

func executeSQL(db *sqlite.Database, schema string) {
	current, _ := os.Getwd()
	parent := filepath.Dir(current)
	seed := parent + "/../../sql/" + schema
	content, err := ioutil.ReadFile(seed)
	if err != nil {
		log.Fatal("Cannot read the seed file: ", seed)
	}
	data := string(content)
	re := regexp.MustCompile(`(\r\n|\n|\r|\t)`)
	sql := re.ReplaceAllString(data, "")
	conn, err := db.GetConnection()
	if err != nil {
		log.Fatal("I could not connect to database so I did not seed.", err)
	}
	_, err = conn.Exec(sql)
	if err != nil {
		log.Fatal("Seed error: ", err)
	}
}

type SQLite struct {
	db *sqlite.Database
}

func NewSQLite(db *sqlite.Database) *SQLite {
	return &SQLite{db: db}
}

func (instance *SQLite) QueryRow(table, i interface{}, id string) *SQLite {
	conn, _ := instance.db.GetConnection()
	query := fmt.Sprintf("select * from %s t where id=?", table)
	err := conn.QueryRowx(query, id).StructScan(i)
	if err != nil {
		log.Fatal("[ERROR] QueryRow:", err)
	}
	return instance
}
