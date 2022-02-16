package storage

import "github.com/fokion/recipeapp/models"

type Storage interface {
	Save(recipe models.Recipe) (*models.Recipe, error)
	Delete(id int64) error
	Update(recipe models.Recipe) error
	Get(id int64) *models.Recipe
	FindByName(name string) (*models.Recipe, error)
}
