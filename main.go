package main 

import (
	"github.com/gin-gonic/gin"

	"github.com/rs/xid"
	"time"
	"net/http"
	// "encoding/json"
	// "io/ioutil"
)
type Recipe struct {
	ID			 string		`json:"id"`
	Name		 string		`json:"name"`
	Tags		 []string	`json:"tags"`
	Ingredients	 []string	`json:"ingredients"`
	Instructions []string	`json:"instructions"`
	PublishedAt	 time.Time	`json:"publishedAt"`
}

var recipes []Recipe 

// Swagger: operation POST
// Create a new recipe
func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe 
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return 
	}
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now() 
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}

// Swagger: operation GET
// List an array of recipes
func ListRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

// Swagger: operation PUT
// Update an existing recipe
func UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	var recipe Recipe  
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
			return 
	}

	index := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index = i 
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found"})
		return 
	}
	recipes[index] = recipe 
	c.JSON(http.StatusOK, recipe)
}

// Swagger: operation DELETE
// Delete an existing recipe
func DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	index := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found"})
		return 
	}

	recipes = append(recipes[:index], recipes[index+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe has been deleted"})
}

func init() {
	recipes = make([]Recipe, 0)

	// file, _ := ioutil.ReadFile("recipes.json")
	// _ = json.Unmarshal([]byte(file), &recipes)

}

func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipesHandler)
	router.PUT("/recipes/:id", UpdateRecipeHandler)
	router.DELETE("recipes/:id", DeleteRecipeHandler)
	router.Run()
}