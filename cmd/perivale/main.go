package main

import (
	"net"

	"github.com/nosborn/go-federation/internal/config"
)

func main() {
	_, err := net.Dial("unix", config.ServerSocketPath)
	if err != nil {
		panic(err)
	}
}
