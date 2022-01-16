package main

import (
	"github.com/Luis-MBL/basic-crud-with-go/database"
	"github.com/Luis-MBL/basic-crud-with-go/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
