package main

import (
	"os"

	runserver "github.com/sajad-dev/go-web-framework/internal/run-server"
	"github.com/sajad-dev/go-web-framework/pkg/connection"
	"github.com/sajad-dev/go-web-framework/pkg/migration"
	"github.com/sajad-dev/go-web-framework/pkg/websocket"

	"github.com/sajad-dev/go-web-framework/pkg/api"
	"github.com/sajad-dev/go-web-framework/pkg/command"
)

func main() {

	runserver.Init()

	connection.Connection()
	if len(os.Args) > 2 {
		command.Handel(os.Args)
		return
	}
	go websocket.Handel()
	migration.Handel()
	api.RouteRun()

	runserver.Run()
}
