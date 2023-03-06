package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgreSqlProvider struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func NewPostgreSqlProvider(userName, password, host, database, port string) PostgreSqlProvider {
	return PostgreSqlProvider{
		Username: userName,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
	}
}

func (s PostgreSqlProvider) ConnectionPool() (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		s.Username,
		s.Password,
		s.Host,
		s.Port,
		s.Database,
	)

	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func (s PostgreSqlProvider) Connection() (*pgxpool.Conn, error) {
	pool, err := s.ConnectionPool()
	if err != nil {
		return nil, err
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
