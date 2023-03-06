package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MySqlProvider struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func NewMySqlProvider(userName, password, host, database, port string) MySqlProvider {
	return MySqlProvider{
		Username: userName,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
	}
}

func (s MySqlProvider) Connection() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		s.Username,
		s.Password,
		s.Host,
		s.Port,
		s.Database,
	)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
