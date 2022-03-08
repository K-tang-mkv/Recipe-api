// swagger:meta

package main

import (
	"context"
	"fmt"
	"log"
	

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	redisStore "github.com/gin-contrib/sessions/redis"
	handlers "github.com/recipes-api/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var recipesHandler *handlers.RecipesHandler 
var authHandler *handlers.AuthHandler

func init() {
	
	ctx := context.Background()
	client, err := mongo.Connect(ctx, 
		options.Client().ApplyURI("mongodb://admin:password@localhost:27017/test?authSource=admin"))

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	
	collection := client.Database("demo").Collection("recipes")
	collectionUsers := client.Database("demo").Collection("users")

	redisClient := redis.NewClient(&redis.Options{
		Addr:		"localhost:6379",
		Password:	"",
		DB:			0,
	})
	status := redisClient.Ping()
	fmt.Println(status)
	
	recipesHandler = handlers.NewRecipesHandler(ctx, collection, redisClient)
	authHandler = handlers.NewAuthHandler(ctx, collectionUsers)
}	


func main() {
	router := gin.Default()

	store, _ := redisStore.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("recipes_api", store))

	router.GET("/recipes", recipesHandler.ListRecipesHandler)

	router.POST("/signin", authHandler.SignInHandler)
	router.POST("/refresh", authHandler.RefreshHandler)
	router.POST("/signout", authHandler.SignOutHandler)

	authorized := router.Group("/")
	authorized.Use(authHandler.AuthMiddleware())
	{
		authorized.POST("/recipes", recipesHandler.NewRecipeHandler)
		authorized.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
		authorized.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
		authorized.GET("/recipes/:id", recipesHandler.GetOneRecipeHandler)
	}

	router.Run()
}