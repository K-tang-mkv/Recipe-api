package routers

import (
	"github.com/gin-contrib/sessions"
	redisStore "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"

	handlers "github.com/recipes-api/handlers"
)

func InitRouter() *gin.Engine {
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

	return router

}
