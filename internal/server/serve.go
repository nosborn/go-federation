package server

import (
	"fmt"
	"net"
	"os"
)

func ListenAndServe(socketPath string) {
	os.Remove(socketPath) // nolint:errcheck

	l, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: listen: %v\n", os.Args[0], err)
		os.Exit(1)
	}
	defer l.Close() // nolint:errcheck

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "accept error: %v\n", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close() // nolint:errcheck
}
