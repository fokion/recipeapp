package storage

import (
	"github.com/fokion/recipeapp/errors"
	"github.com/fokion/recipeapp/models"
	"strings"
)

type MemoryStorage struct {
	recipes []models.Recipe
}

func (s *MemoryStorage) Save(recipe models.Recipe) (*models.Recipe, error) {
	recipeFound, err := s.FindByName(recipe.Name)
	if err != nil {
		return nil, err
	}
	if recipeFound != nil {
		return nil, errors.APIError{Detail: "Recipe exists already", Service: "MemoryStorage", Code: errors.REC_01}
	}
	recipe.ID = int64(len(s.recipes) + 1)
	s.recipes = append(s.recipes, recipe)
	return &recipe, nil
}

func (s *MemoryStorage) Delete(id int64) error {
	return nil
}
func (s *MemoryStorage) Update(recipe models.Recipe) error {
	return nil
}
func (s *MemoryStorage) Get(id int64) *models.Recipe {
	var recipe *models.Recipe
	for i := 0; i < len(s.recipes); i++ {
		recipe = &s.recipes[i]
		if recipe.ID == id {
			return recipe
		}
	}
	return nil
}

func (s *MemoryStorage) FindByName(name string) (*models.Recipe, error) {
	var recipe *models.Recipe
	for i := 0; i < len(s.recipes); i++ {
		recipe = &s.recipes[i]
		if strings.EqualFold(name, recipe.Name) {
			return recipe, nil
		}
	}
	return recipe, nil
}
