package application

import (
	"api-key-middleware/internal/core/ports"
	"api-key-middleware/internal/core/services"
	"api-key-middleware/internal/handlers"
	"api-key-middleware/internal/repositories"
	"api-key-middleware/pkg/cache"
	"api-key-middleware/pkg/database"
	"api-key-middleware/pkg/validator"
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

type Dependencies struct {
	cache ports.Cache
	mysql *sql.DB

	healthChecksHandlers ports.HealthChecksHandlers
	userHandlers         ports.UserHandlers
}

func NewDependencies() Dependencies {
	cacheInstance := NewCacheConnection()
	dbConnection := NewPostgreSQLDBConnection()
	reqValidator := NewValidator()

	return NewDependenciesWith(cacheInstance, dbConnection, reqValidator)
}

func NewDependenciesWith(cacheInstance ports.Cache, dbConnection *sql.DB, reqValidator ports.Validator) Dependencies {
	return Dependencies{
		cache:                cacheInstance,
		mysql:                dbConnection,
		healthChecksHandlers: newHealthChecksHandlers(cacheInstance, dbConnection),
		userHandlers:         newUserHandlers(dbConnection, reqValidator),
	}
}

func newHealthChecksHandlers(cache ports.Cache, dbConnection *sql.DB) ports.HealthChecksHandlers {
	return handlers.NewHealthChecksHandlers(cache, dbConnection)
}

func newUserHandlers(dbConnection *sql.DB, validator ports.Validator) ports.UserHandlers {
	repo := repositories.NewUserRepository(dbConnection)
	userService := services.NewUserService(repo)
	return handlers.NewUserHandlers(userService, validator)
}

func NewPostgreSQLDBConnection() *sql.DB {
	time.Sleep(time.Second * 3)
	dbConnection, err := database.NewMySqlProvider(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("MYSQL_DATABASE"),
		os.Getenv("DB_PORT"),
	).Connection()
	if err != nil {
		log.Fatalf("connection error: %v \n", err)
	}

	return dbConnection
}

func NewValidator() ports.Validator {
	return validator.NewValidator()
}

func NewCacheConnection() *cache.Cache {
	cacheHost := fmt.Sprintf("%s:%s", os.Getenv("CACHE_HOST"), os.Getenv("CACHE_PORT"))
	return cache.NewCache(cacheHost)
}
