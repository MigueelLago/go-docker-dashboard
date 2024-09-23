package main

import (
	"github.com/MigueelLago/go-docker-dashboard/routes"
	"github.com/docker/docker/client"
)

func main() {
	// Init client Docker
	apliClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	defer apliClient.Close()

	// Init App and Routes
	routes.Initialize(apliClient)
}
