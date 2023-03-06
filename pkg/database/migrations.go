package database

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/clickhouse"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DBDriverName string

const (
	Mysql      DBDriverName = "mysql"
	Postgres                = "postgres"
	Clickhouse              = "clickhouse"
)

var unknownDBDriver = errors.New("error: unknown database driver name")

const migrationFilesPath = "database/migrations"

type MigratorContract interface {
	New() error
	Up() error
	Down() error
	Version() (version uint, dirty bool, err error)
	Force(version int) error
}

var _ MigratorContract = (*Migrator)(nil)

type Migrator struct {
	migrate *migrate.Migrate
}

func NewMigrator() (Migrator, error) {
	migrator, err := migrate.New(
		fmt.Sprintf("file://%s", migrationFilesPath),
		getDatabaseDSN(DBDriverName(os.Getenv("DB_DRIVER"))),
	)
	if err != nil {
		return Migrator{}, fmt.Errorf("instance error: %v \n", err)
	}

	return Migrator{migrate: migrator}, nil
}

func NewMigratorWithMigrationsPath(path string) (Migrator, error) {
	migrator, err := migrate.New(
		fmt.Sprintf("file://%s", path),
		getDatabaseDSN(DBDriverName(os.Getenv("DB_DRIVER"))),
	)
	if err != nil {
		return Migrator{}, fmt.Errorf("instance error: %v \n", err)
	}

	return Migrator{migrate: migrator}, nil
}

func getDatabaseDSN(driverName DBDriverName) string {
	var dsn string
	switch driverName {
	case Mysql:
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"),
		)
	case Clickhouse:
		dsn = fmt.Sprintf(
			"clickhouse://%s:%s?username=%s&password=%s&database=%s&x-multi-statement=true&x-migrations-table=migration_table&debug=true",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DATABASE"),
		)
	case Postgres:
		dsn = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"),
		)
	default:
		log.Fatalln(unknownDBDriver)
	}
	return dsn
}

func (m *Migrator) New() error {
	reader := bufio.NewReader(os.Stdin)
	println("Enter new migration name:")

	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	timePoint := time.Now().Unix()
	upFileName := fmt.Sprintf("%s/%d_create_%s_table.up.sql", migrationFilesPath, timePoint, text)
	downFileName := fmt.Sprintf("%s/%d_drop_%s_table.down.sql", migrationFilesPath, timePoint, text)

	emptyFile, err := os.Create(upFileName)
	if err != nil {
		return err
	}

	err = emptyFile.Close()
	if err != nil {
		return err
	}

	emptyFile, err = os.Create(downFileName)
	if err != nil {
		return err
	}

	err = emptyFile.Close()
	if err != nil {
		return err
	}

	log.Println(fmt.Sprintf("Migrations (%s, %s) created", upFileName, downFileName))

	return nil
}

func (m *Migrator) Up() error {
	return m.migrate.Up()
}

func (m *Migrator) Down() error {
	return m.migrate.Down()
}

func (m *Migrator) Version() (version uint, dirty bool, err error) {
	return m.migrate.Version()
}

func (m *Migrator) Force(version int) error {
	return m.migrate.Force(version)
}
