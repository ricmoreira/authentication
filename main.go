package main

import (
	"authentication/containers"
	"authentication/server"
)

func main() {

	container := containers.BuildContainer()

	err := container.Invoke(func(server *server.Server) {
		server.Run()
	})

	if err != nil {
		panic(err)
	}
}
