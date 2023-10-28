package main

import "massengerGrpc/internal/server"

func main() {
	s := server.New()
	s.Start()
}
