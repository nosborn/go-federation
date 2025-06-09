package main

import (
	"github.com/nosborn/go-federation/internal/config"
	"github.com/nosborn/go-federation/internal/server"
)

func main() {
	server.ListenAndServe(config.ServerSocketPath)
}
