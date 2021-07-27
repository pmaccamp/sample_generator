package models

import "fmt"

type Meal struct {
	Number int
	Foods []Food
}

func (m Meal) GetCalories() float64 {
	cals := 0.0
	for _, food := range m.Foods {
		cals += food.Calories
	}

	return cals
}

func (m Meal) Print() {
	fmt.Printf("Meal %v \n", m.Number)

	for _, food := range m.Foods {
		fmt.Printf("    %v\n", food)
	}
}