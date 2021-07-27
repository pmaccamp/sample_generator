package generator

import (
	"sample_generator/models"
)

type Generator interface {
	Generate(potentialFoods []models.Food, exclusions []string, calorieTarget float64) (*models.Day, error)
}
