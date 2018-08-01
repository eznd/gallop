package main

import (
	"os"

	"eznd/gallop/server"
)

func main() {
	rootPath := os.Getenv("GALLOP_ROOT")
	if rootPath == "" {
		rootPath = "/go/src/eznd/gallop/"
	}

	host := os.Getenv("GALLOP_HOST")

	port := os.Getenv("GALLOP_PORT")
	if port == "" {
		port = "9101"
	}

	srv := server.New(rootPath, host, port)
	srv.StartRouter()
}
