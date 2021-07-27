package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"sample_generator/generator"
	"sample_generator/models"
	"strings"
)

func getGenerator() generator.Generator {
	// TODO Replace me with an actual generator
	return generator.MockGenerator{}
}

func loadFoods() []models.Food {
	var foods []models.Food
	foodsJson, _ := ioutil.ReadFile("foods.json") // filename is the JSON file to read
	err := json.Unmarshal(foodsJson, &foods)
	if err != nil {
		panic(err)
	}
	return foods
}

func main() {
	calorieTargetPtr := flag.Float64("cals", 2000, "calorie target for the day")
	exclusionsStr := flag.String("excls", "", "comma separated list of exclusion keywords")
	flag.Parse()

	if calorieTargetPtr == nil {
		panic("invalid calorie target")
	}
	calorieTarget := *calorieTargetPtr
	var exclusions []string
	if *exclusionsStr != "" {
		exclusions = strings.Split(*exclusionsStr, ",")
	}

	fmt.Printf("Generating a day with %v calories target and [%v] exclusions\n", calorieTarget, strings.Join(exclusions, ","))
	day, err := getGenerator().Generate(loadFoods(), exclusions, calorieTarget)
	if err != nil {
		panic(err)
	}
	day.Print()
	day.PrintError(calorieTarget)
}
