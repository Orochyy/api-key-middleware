package repositories

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"time"
)

var (
	builderPSQL = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
)

type Informer interface {
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

// dst should be a slice of pointers to data
func queryRow(conn Informer, query string, args []interface{}, dst []interface{}) error {
	ctx := context.Background()
	if err := conn.QueryRow(ctx, query, args...).Scan(dst...); err != nil {
		return fmt.Errorf("error scanning result: %w", err)
	}
	return nil
}

func puint(id uint) *uint {
	return &id
}

func timeNow() time.Time {
	return time.Now().UTC()
}
