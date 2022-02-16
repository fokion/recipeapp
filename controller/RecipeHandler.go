package controller

import (
	"fmt"
	"github.com/fokion/recipeapp/errors"
	"github.com/fokion/recipeapp/models"
	"github.com/fokion/recipeapp/storage"
	_ "github.com/fokion/recipeapp/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

type request struct {
	id int64
}

type RecipeHandler interface {
	SaveRecipe(c *gin.Context)
	GetRecipe(c *gin.Context)
}

type recipeController struct {
	StorageService storage.Storage
}

func RecipeController(storage storage.Storage) RecipeHandler {
	return &recipeController{
		storage,
	}
}

func (h *recipeController) SaveRecipe(c *gin.Context) {
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.APIError{Service: "RecipeHandler", Detail: "Malformed JSON", Code: errors.REC_02}})
		return
	}

	pointerToRecipe, saveErr := h.StorageService.Save(recipe)
	if saveErr == nil {
		c.JSON(http.StatusOK, &pointerToRecipe)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": saveErr})
	return
}
func (h *recipeController) GetRecipe(c *gin.Context) {
	var req request
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"msg": errors.APIError{Service: "RecipeHandler", Detail: err.Error(), Code: errors.REC_03}})
		return
	}
	recipePointer := h.StorageService.Get(req.id)
	if recipePointer == nil {
		c.JSON(400, gin.H{"msg": errors.APIError{Service: "RecipeHandler", Detail: fmt.Sprintf("Could not find this recipe with id %d", req.id), Code: errors.REC_04}})
		return
	}
	c.JSON(200, &recipePointer)
}
