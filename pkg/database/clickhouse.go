package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ClickHouse/clickhouse-go"
	log "github.com/sirupsen/logrus"
	"time"
)

type ClickHouse struct {
	connection *sql.DB
}

func NewClickhouseConnection(host, port, user, password, database string) (ClickHouse, error) {
	dataSourceName := fmt.Sprintf(
		"tcp://%s:%s?username=%s&password=%s&database=%s&x-multi-statement=true&debug=true",
		host,
		port,
		user,
		password,
		database,
	)

TODO:
	connection, err := sql.Open("clickhouse", dataSourceName)
	if err != nil {
		log.Printf("Try to reconnect to %s", dataSourceName)
		time.Sleep(time.Second * 10)
		goto TODO
	}
	if err := connection.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			return ClickHouse{}, errors.New(
				fmt.Sprintf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace),
			)
		} else {
			return ClickHouse{}, err
		}
	}

	return ClickHouse{
		connection: connection,
	}, nil
}

func (c ClickHouse) Connection() *sql.DB {
	return c.connection
}

func (c ClickHouse) Exec(query string) (sql.Result, error) {
	return c.connection.Exec(query)
}

func (c ClickHouse) Close() error {
	return c.connection.Close()
}

func (c ClickHouse) Begin() (*sql.Tx, error) {
	tx, err := c.connection.Begin()
	if err != nil {
		return nil, err
	}

	return tx, nil
}
