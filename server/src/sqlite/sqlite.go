package sqlite

import (
	"fmt"
	"log"
	"server/libs/filter"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	connection *sqlx.DB
	tx         *sqlx.Tx
	config     Config
}

type Config struct {
	URL string
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

func (db *Database) Get(query string, found interface{}, args ...interface{}) error {
	conn, err := db.GetConnection()
	if err != nil {
		return err
	}
	return conn.Get(found, query, args...)
}

func (db *Database) QueryRow(query string, found interface{}, args ...interface{}) error {
	conn, err := db.GetConnection()
	if err != nil {
		return err
	}
	err = conn.QueryRowx(query, args...).StructScan(found)
	return err
}

func (db *Database) Exec(query string, args ...interface{}) error {
	conn, err := db.GetConnection()
	if err != nil {
		return err
	}
	_, err = conn.Exec(query, args...)
	return err
}

func (db *Database) ExecAndGetLastID(query string, args ...interface{}) (string, error) {
	conn, err := db.GetConnection()
	if err != nil {
		return "", err
	}
	result, err := conn.Exec(query, args...)
	if err != nil {
		return "", err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", id), nil
}

func (db *Database) QueryAll(query string, found interface{}, args ...interface{}) error {
	conn, err := db.GetConnection()
	if err != nil {
		return err
	}
	err = conn.Select(found, query, args...)
	return err
}

type ComplexQuery struct {
	pool       *sqlx.DB
	Filters    filter.Request
	Table      string
	Total      *int
	fields     string
	tables     string
	groupBy    string
	joins      string
	countQuery string
}

func (db *Database) Select() *ComplexQuery {
	pool, _ := db.GetConnection()
	return &ComplexQuery{pool: pool}
}

func (cq *ComplexQuery) WithFields(fields []string) *ComplexQuery {
	cq.fields = strings.Join(fields, ", ")
	return cq
}

func (cq *ComplexQuery) FromTables(tables []string) *ComplexQuery {
	cq.tables = strings.Join(tables, ", ")
	return cq
}

func (cq *ComplexQuery) WithJoins(joins []string) *ComplexQuery {
	cq.joins = strings.Join(joins, " AND ")
	return cq
}

func (cq *ComplexQuery) From(table string) *ComplexQuery {
	cq.Table = table
	return cq
}

func (cq *ComplexQuery) Applying(filters filter.Request) *ComplexQuery {
	cq.Filters = filters
	return cq
}

func (cq *ComplexQuery) GroupBy(field string) *ComplexQuery {
	cq.groupBy = field
	return cq
}

func (cq *ComplexQuery) WithCount(total *int, table string) *ComplexQuery {
	cq.Total = total
	cq.countQuery = "SELECT COUNT(*) FROM " + table
	return cq
}

func (cq *ComplexQuery) WithTotal(total *int) *ComplexQuery {
	cq.Total = total
	return cq
}

func (cq *ComplexQuery) ForEach(row interface{}, f func(i interface{})) error {
	conn := cq.pool
	if conn == nil {
		return fmt.Errorf("Database connection lost")
	}
	table := cq.Table
	if cq.tables != "" {
		table = cq.tables
	}
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s %s", table, cq.Filters.Where())
	if cq.countQuery != "" {
		query = fmt.Sprintf("%s %s", cq.countQuery, cq.Filters.Where())
	}
	err := conn.Get(cq.Total, query)
	if err != nil {
		return err
	}
	fields := "*"
	if cq.fields != "" {
		fields = cq.fields
	}
	sql := cq.Filters.Where()
	if cq.joins != "" {
		sql += "AND " + cq.joins
	}
	sql += cq.Filters.GroupBy()
	sql += cq.Filters.OrderBy()
	sql += cq.Filters.Limit()
	query = fmt.Sprintf("SELECT %s FROM %s %s", fields, table, sql)
	rows, err := conn.Queryx(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.StructScan(row)
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
