package main

import (
	"os"

	"github.com/yucacodes/secure-port-forwarding/client"
	"github.com/yucacodes/secure-port-forwarding/server"
)

func main() {

	args := os.Args[1:]
	if len(args) < 1 {
		return
	}

	module := args[0]
	if module == "server" {
		server.Main()
	} else if module == "client" {
		client.Main()
	}
}
