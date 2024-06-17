package main

import (
	"github.com/1Nelsonel/Savannah/database"
	"github.com/1Nelsonel/Savannah/routes"
	"github.com/gin-gonic/gin"

)

// Middleware to initialize db connections
func init() {
	database.ConnectDB()
}

func main() {
	r := gin.Default()

	// routes
	routes.SetupRouter(r)

	r.Run(":3000")
}