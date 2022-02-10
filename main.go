package main

//importing neccessary libararies
import (
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"

	"Project2/dependencies"
	"Project2/log"
	"Project2/server"
)

func main() {
	log.Setup()
	dependency := dependencies.InitDependency()
	router := server.Router(dependency)
	startApp(router)
}

func startApp(router *mux.Router) {
	app := cli.NewApp()
	app.Name = "Project2"
	app.Version = "0.0.1"
	app.Commands = []*cli.Command{
		{
			Name:        "start-app",
			Description: "this starts the web app",
			Action: func(context *cli.Context) error {
				server.StartServer(router)
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
