package main

import (
	"log"

	"bitbucket.org/topdoctors/tools/controllers"

	"basic-go-training/internal/config"
	"basic-go-training/internal/controllers/health"
	"basic-go-training/internal/controllers/patients"
	"basic-go-training/internal/database"
)

func main() {
	if err := config.ParseSettings(); err != nil {
		log.Fatal(err)
	}

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	server := controllers.NewServer(
		controllers.WithCors(),
		controllers.WithSwagger(),
	)
	health.NewServer(server.Engine())
	patients.NewServer(server.Engine())

	log.Fatal(server.Run(":9099"))
}
