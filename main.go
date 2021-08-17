package main

import "github.com/CavasCallahan/firstGo/server"

func main() {
	server := server.NewServer()
	server.Run()
}
