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
	"github.com/gin-contrib/sessions"
	redisStore "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"

	handlers "github.com/recipes-api/handlers"
)

func main() {
	router := gin.Default()

	store, _ := redisStore.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("recipes_api", store))

	router.GET("/recipes", handlers.RecipesHandler.ListRecipesHandler)

	router.POST("/signin", handlers.AuthHandler.SignInHandler)
	router.POST("/refresh", handlers.AuthHandler.RefreshHandler)
	router.POST("/signout", handlers.AuthHandler.SignOutHandler)

	authorized := router.Group("/")
	authorized.Use(handlers.AuthHandler.AuthMiddleware())
	{
		authorized.POST("/recipes", handlers.RecipesHandler.NewRecipeHandler)
		authorized.PUT("/recipes/:id", handlers.RecipesHandler.UpdateRecipeHandler)
		authorized.DELETE("/recipes/:id", handlers.RecipesHandler.DeleteRecipeHandler)
		authorized.GET("/recipes/:id", handlers.RecipesHandler.GetOneRecipeHandler)
	}

	router.Run()
}
