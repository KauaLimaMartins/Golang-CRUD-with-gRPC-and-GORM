package main

import (
	"github.com/kaualimamartins/gRPC-Go-Gorm/customers"
	"github.com/kaualimamartins/gRPC-Go-Gorm/database"
)

func main() {
	// Connect with database
	database.ConnectWithDatabase()

	// Create gRPC server
	customers.CreateGRPCServer()
}
