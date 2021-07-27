package models

import "fmt"

type Food struct {
	Name string
	Calories float64
}

func (f Food) String() string {
	return fmt.Sprintf("%v - %v cals", f.Name, f.Calories)
}