package main

import (
	"log"

	"bitbucket.org/topdoctors/tools/controllers"

	"basic-go-training/internal/controllers/health"
)

func main() {
	server := controllers.NewServer(
		controllers.WithCors(),
		controllers.WithSwagger(),
	)
	health.NewServer(server.Engine())

	log.Fatal(server.Run(":9099"))
}
