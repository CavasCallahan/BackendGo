package main

import (
	"github.com/CavasCallahan/firstGo/server"
	"github.com/CavasCallahan/firstGo/server/database"
)

func main() {
	database.StartDB() // Start's the database connection

	server := server.NewServer() //Create's a new server
	server.Run()                 //Start's the new server
}
