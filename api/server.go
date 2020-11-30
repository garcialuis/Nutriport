package api

import (
	"os"

	"github.com/garcialuis/Nutriport/api/controllers"
)

var server = controllers.Server{}

func Run() {

	DbName := os.Getenv("DB_NAME")
	server.Initialize(os.Getenv("DB_POSTGRES_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), DbName)
	// Using Load from seed is used for 1st use and/or during development
	// seed.Load(server.DB)
	server.Run(":8085")
}
