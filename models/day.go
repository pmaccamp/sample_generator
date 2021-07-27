package models

import (
	"fmt"
	"math"
)

type Day struct {
	Meals []Meal
}

func (d Day) GetCalories() float64 {
	cals := 0.0
	for _, meal := range d.Meals {
		cals += meal.GetCalories()
	}

	return cals
}

func (d Day) Print() {
	fmt.Printf("Day %v cals\n", d.GetCalories())
	for _, meal := range d.Meals {
		meal.Print()
	}
}

func (d Day) PrintError(target float64) {
	dayCals := d.GetCalories()
	fmt.Printf("Calorie error: %v %v%%\n", dayCals-target, int(100*math.Abs(dayCals-target)/target))
}
