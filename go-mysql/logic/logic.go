package logic

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type Logic struct {
	db *sql.DB
}

func NewLogic(db *sql.DB) *Logic {
	return &Logic{
		db: db,
	}
}

func (l *Logic) Select(fields []string, table string) (*sql.Rows, error) {
	query := fmt.Sprintf("select %s from %s", strings.Join(fields, ","), table)
	log.Println(query)
	return l.db.Query(query)
}

func (l *Logic) SelectWithWhere(fields []string, table, conditions string, args ...any) (*sql.Rows, error) {
	query := fmt.Sprintf("select %s from %s where %s", strings.Join(fields, ","), table, conditions)
	log.Println(query, args)
	return l.db.Query(query, args...)
}
