// Recipes API
//
// This is a sample recipes API. You can find out more about the API at https://github.com/PacktPublishing/Building-Distributed-Applications-in-Gin.
//
//	Schemes: http
//  Host: localhost:8080
//	BasePath: /
//	Version: 1.0.0
//	Contact: nowhere
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
// swagger:meta

package main

import (
	"github.com/recipes-api/routers"
)

func main() {

	router := routers.InitRouter()

	router.Run()
}
