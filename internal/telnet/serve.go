package telnet

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"os/exec"
)

func ListenAndServe() {
	listener, err := net.Listen("tcp", ":23")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: listen: %v\n", os.Args[0], err)
		os.Exit(1)
	}

	serve(listener)
}

func ListenAndServeTLS(certFile string, keyFile string) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if nil != err {
		panic(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", ":992", config)
	if nil != err {
		fmt.Fprintf(os.Stderr, "%s: listen: %v\n", os.Args[0], err)
		os.Exit(1)
	}

	serve(listener)
}

func serve(listener net.Listener) {
	defer listener.Close() // nolint:errcheck

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "accept error: %v\n", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close() // nolint:errcheck

	cmd := exec.Command("/app/bin/login")

	err := cmd.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: child finished with error: %v", os.Args[0], err)
	}
}
