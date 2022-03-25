package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect(psqlconn string) (*sql.DB, error) {
	pg, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, errorFmt(err)
	}

	err = pg.Ping()
	if err != nil {
		return nil, errorFmt(err)
	}

	return pg, nil
}

func Disconnect(pg *sql.DB) error {
	err := pg.Close()
	return err
}

func Query(pg *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	result, err := pg.Query(query, args...)
	return result, err
}

func QueryRow(pg *sql.DB, query string, args ...interface{}) *sql.Row {
	result := pg.QueryRow(query, args...)
	return result
}

func ScanRow(row *sql.Row, dest ...interface{}) error {
	err := row.Scan(dest...)
	return err
}

func ScanRows(rows *sql.Rows, dest ...interface{}) error {
	err := rows.Scan(dest...)
	return err
}

func errorFmt(err error) error {
	return fmt.Errorf("Falha ao conectar com o banco de dados: %s", err)
}
