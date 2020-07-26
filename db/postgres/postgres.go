package postgres

import (
	"fmt"
	"github.com/jackc/pgx"
)

const SqlTimeLayout = "02-01-06 15:04:05"

// Пул подключений к БД. Не забывать закрыть пул после использования
func NewConnPool(login string, password string, host string, port int, name string) (pool *pgx.ConnPool, err error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", login, password, host, port, name)
	var cfg pgx.ConnConfig
	cfg, err = pgx.ParseConnectionString(connStr)

	if err != nil {
		return
	}
	pool, err = pgx.NewConnPool(pgx.ConnPoolConfig{MaxConnections: 2, ConnConfig: cfg})

	if err != nil {
		return
	}
	return
}
