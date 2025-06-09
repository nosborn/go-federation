package main

import (
	"os"
	"syscall"
)

func main() {
	args := []string{"perivale"}
	env := os.Environ()
	err := syscall.Exec("/app/bin/perivale", args, env)
	if err != nil {
		panic(err)
	}
}
