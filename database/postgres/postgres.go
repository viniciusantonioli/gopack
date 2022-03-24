package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func Connect(pgConnString string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), pgConnString)
	if err != nil {
		return nil, fmt.Errorf("Não foi possível conectar a base de dados: %v\n", err)
	}
	return conn, nil
}

func Query(conn *pgx.Conn, query string, args ...interface{}) (pgx.Rows, error) {
	rows, err := conn.Query(context.Background(), query, args...)
	if err != nil {
		return nil, fmt.Errorf("Não foi possível executar a query: %v\n", err)
	}
	return rows, nil
}

func Disconnect(conn *pgx.Conn) error {
	err := conn.Close(context.Background())
	if err != nil {
		return fmt.Errorf("Não foi possível desconectar da base de dados: %v\n", err)
	}
	return nil
}
