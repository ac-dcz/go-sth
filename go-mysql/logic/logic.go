package logic

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type TxItem struct {
	Sql  string
	Vars []any
}

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

func (l *Logic) Txn(items []TxItem) (err error) {
	tx, err := l.db.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()
	for i, item := range items {
		result, err := tx.Exec(item.Sql, item.Vars...)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println("=====================================")
		fmt.Printf("Query %d: %s %v\n", i, item.Sql, item.Vars)
		n, _ := result.RowsAffected()
		fmt.Printf("Results %d: Affected rows %d\n", i, n)
	}
	return nil
}
