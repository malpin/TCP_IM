package main

import "TCP_IM/server"

func main() {
	newServer := server.NewServer("0.0.0.0", 2020)
	newServer.Start()
}
