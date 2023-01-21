package mysql

import (
	"database/sql"

	gomysql "github.com/go-sql-driver/mysql"
)

func New(config *Config) (*sql.DB, error) {
	cfg := gomysql.Config{
		User:   config.User,
		Passwd: config.Password,
		Addr:   config.Address,
		DBName: config.DB,
	}
	dsn := cfg.FormatDSN()
	return sql.Open("mysql", dsn)
}
