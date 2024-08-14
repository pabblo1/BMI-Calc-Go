package main

import (
	"fmt"
	"math"
)

func main() {

	var mass float64
	var height float64

	fmt.Println("Enter mass (in kg): ")
	fmt.Scanln(&mass)
	fmt.Println("Enter height (in m):")
	fmt.Scanln(&height)

	var bmi float64 = mass / math.Pow(height, 2)

	if bmi < 18.5 {
		fmt.Println("You are underweight!")
	} else if bmi >= 18.5 && bmi < 25 {
		fmt.Println("You are healthy!")
	} else if bmi > 25 {
		fmt.Println("You are overweight!")
	} else {
		fmt.Println("An error occurred")
	}

}
