package main

import (
	"api-key-middleware/internal/application"
)

// @title API Key Middleware
// @version 1.0
func main() {
	app := application.NewApp()
	app.RunServer()
	app.CleanupTasks()

}
