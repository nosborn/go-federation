package main

import (
	"github.com/nosborn/go-federation/internal/telnet"
)

func main() {
	go telnet.ListenAndServe()
	telnet.ListenAndServeTLS("/app/ssl/telnetd.crt", "/app/ssl/telnetd.key")
}
