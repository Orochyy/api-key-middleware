package application

import (
	"api-key-middleware/internal/middleware"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Router(server *fiber.App, dependencies Dependencies) {
	server.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	server.Use(logger.New(logger.Config{
		Format:     "[${time}] [${ip}]:${port} ${method} ${path} ${status} pid=${pid}\n",
		TimeFormat: DateTimeLayout,
	}))

	docs := server.Group("/docs")
	docs.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "sqr",
		},
	}))
	docs.Get("/*", swagger.HandlerDefault)

	server.Get("/health_checks", dependencies.healthChecksHandlers.HealthChecks)

	server.Use(middleware.AuthMiddleware(dependencies.mysql))

	v1 := server.Group("/api/v1")

	Routes := v1.Group("/")

	userRoutes := Routes.Group("user")
	userRoutes.Get("/profile", dependencies.userHandlers.Get)

}
