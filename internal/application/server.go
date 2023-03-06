package application

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
)

const (
	DateTimeLayout = "15:04:05 02-01-2006"
)

type HttpServer struct {
	client *Client
}

func NewHttpServer(options ...Option) *HttpServer {
	client := &Client{
		host: defaultHttpServerHost,
		port: defaultHttpServerPort,
	}
	for _, option := range options {
		option(client)
	}

	return &HttpServer{
		client: client,
	}
}

func (s *HttpServer) GetServer() *fiber.App {
	prefork, err := strconv.ParseBool(os.Getenv("SERVER_PREFORK"))
	if err != nil {
		prefork = false
	}

	return fiber.New(fiber.Config{
		DisableStartupMessage: false,
		Prefork:               prefork,
		CaseSensitive:         false,
		StrictRouting:         true,
		IdleTimeout:           idleTimeout,
		ServerHeader:          "CustomServer",
		JSONEncoder:           json.Marshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			log.Info(err.Error())
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

			if code == http.StatusInternalServerError {
				return ctx.Status(code).SendString(http.StatusText(http.StatusInternalServerError))
			}

			return ctx.Status(code).SendString(err.Error())
		},
	})
}

func (s *HttpServer) GetDSN() string {
	return fmt.Sprintf("%s:%s", s.client.host, s.client.port)
}
