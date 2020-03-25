package sqlite

import (
	"fmt"
	"log"
	"server/libs/filter"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	connection *sqlx.DB
	tx         *sqlx.Tx
    config Config
}

type Config struct {
	URL             string
}

func New(config Config) *Database {
    return &Database{config: config}
}

func (db *Database) Connect() (*sqlx.DB, error) {
	conn, err := sqlx.Connect("sqlite3", db.config.URL)
	if err != nil {
		log.Println("[ERROR]", err)
		return nil, fmt.Errorf("Nosso banco de dados está em manutenção.")
	}
	return conn, nil
}

func (db *Database) GetConnection() (*sqlx.DB, error) {
	var err error
	if db.connection == nil {
		db.connection, err = db.Connect()
	}
	return db.connection, err
}

func (db *Database) QueryRow(query string, found interface{}, args ...interface{}) error {
	conn, err := db.GetConnection()
	if err != nil {
		return err
	}
	err = conn.QueryRowx(query, args...).StructScan(&found)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) Exec(query string, args ...interface{}) error {
	conn, err := db.GetConnection()
	if err != nil {
		return err
	}
	_, err = conn.Exec(query, args...)
    return err
}

type ComplexQuery struct {
	pool    *sqlx.DB
	Filters filter.Request
	Table   string
	Total   *int
}

func (db *Database) Select() *ComplexQuery {
	pool, _ := db.GetConnection()
	return &ComplexQuery{pool: pool}
}

func (cq *ComplexQuery) From(table string) *ComplexQuery {
	cq.Table = table
	return cq
}

func (cq *ComplexQuery) Applying(filters filter.Request) *ComplexQuery {
	cq.Filters = filters
	return cq
}

func (cq *ComplexQuery) WithTotal(total *int) *ComplexQuery {
	cq.Total = total
	return cq
}

func (cq *ComplexQuery) ForEach(f func(i interface{})) error {
	conn := cq.pool
	if conn == nil {
		return fmt.Errorf("Database connection lost")
	}
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s %s", cq.Table, cq.Filters.Where())
	err := conn.QueryRowx(query).StructScan(&cq.Total)
	if err != nil {
		return err
	}
	table := cq.Table
	query = fmt.Sprintf("SELECT * FROM %s %s", table, cq.Filters.SQL())
	rows, err := conn.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var row interface{}
		err = rows.Scan(&row)
		if err != nil {
			return err
		}
		f(row)
	}
	if rows.Err() != nil {
		return fmt.Errorf("%v", rows.Err())
	}
	return nil
}
