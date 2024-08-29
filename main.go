package main

import (
	"log"
	"task/config"
	"task/database"
	"task/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// loading config files such as .env and other configs
	config.LoadConfig()

	// create the initial connection to the db
	database.Connect()

	// migration to set the schema in the db
	database.Migrate()

	// initilizing fiber application
	app := fiber.New()

	// set up end points for the application
	routes.SetupRoutes(app)

	// run the applicatoin on port 3000
	log.Fatal(app.Listen(":3000"))

}
