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

// Init the mux router
/*router := mux.NewRouter()

// Route handles & endpoints

// Get all movies
router.HandleFunc("/getusers/", modules.GetUser).Methods("GET")

// Create a movie
router.HandleFunc("/createusers/", modules.CreateUsers).Methods("POST")

//Delete a user
router.HandleFunc("/deleteuser/{email}", modules.DeleteUser).Methods("DELETE")

//Update a user
router.HandleFunc("/updateuser/{email}", modules.UpdateUser).Methods("PUT")

fmt.Println("Server at 8080")
log.Fatal(http.ListenAndServe(":8000", router))*/

/*
import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file://db/migrations",
		"postgres://postgres:0712@localhost:5432/test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

}
*/
