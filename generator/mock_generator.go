package generator

import "sample_generator/models"

type MockGenerator struct {
}

func (g MockGenerator) Generate(potentialFoods []models.Food, exclusions []string, calorieTarget float64) (*models.Day, error) {
	//returns dummy Day
	return &models.Day{
		Meals: []models.Meal{
			{
				Number: 1,
				Foods: []models.Food{
					{
						Name:     "Eggs",
						Calories: 500,
					},
				},
			},
			{
				Number: 2,
				Foods: []models.Food{
					{
						Name:     "Sandwich",
						Calories: 500,
					},
				},
			},
			{
				Number: 3,
				Foods: []models.Food{
					{
						Name:     "Steak",
						Calories: 500,
					},
				},
			},
		},
	}, nil
}
