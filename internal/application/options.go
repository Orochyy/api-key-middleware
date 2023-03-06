package application

import "time"

const (
	defaultHttpServerHost = "0.0.0.0"
	defaultHttpServerPort = "8080"
	idleTimeout           = 5 * time.Second
)

type Client struct {
	host, port string
}

type Option func(*Client)

func WithPort(port string) Option {
	return func(client *Client) {
		client.port = port
	}
}

func WithHost(host string) Option {
	return func(client *Client) {
		client.host = host
	}
}
