package api

import (
	"os"

	"github.com/garcialuis/Nutriport/api/controllers"
)

var server = controllers.Server{}

func Run() {

	// TODO: Check environment value to determin weather to get db_name from env
	//   and wether to run seed.Load() or not.
	DB_NAME := os.Getenv("DB_NAME")
	DB_NAME = "nutriport" // This line is used for dev purposes
	server.Initialize(os.Getenv("DB_POSTGRES_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), DB_NAME)
	// Using Load from seed is used for 1st use and/or during development
	// seed.Load(server.DB)
	server.Run(":8085")
}
