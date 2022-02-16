package main

import (
	"github.com/fokion/recipeapp/controller"
	"github.com/fokion/recipeapp/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	recipeController := controller.RecipeController(&storage.MemoryStorage{})
	router.GET("/liveliness", func(context *gin.Context) {
		context.JSON(200, gin.H{})
	})
	router.POST("/recipe", recipeController.SaveRecipe)
	router.GET("/recipe/:id", recipeController.GetRecipe)
	router.Run(":8080")
}
