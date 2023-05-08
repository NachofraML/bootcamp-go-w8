package main

import (
	"fmt"
)

func main() {
	const (
		dog    = "dog"
		cat    = "cat"
		mouse  = "mouse"
		spider = "spider"
	)

	animalDog, msg := animal(dog)
	if msg != "" {
		fmt.Println(msg)
	}

	animalCat, msg := animal(cat)
	if msg != "" {
		fmt.Println(msg)
	}

	animalMouse, msg := animal(mouse)
	if msg != "" {
		fmt.Println(msg)
	}

	animalSpider, msg := animal(spider)
	if msg != "" {
		fmt.Println(msg)
	}

	var amount float64
	amount += animalDog(10)
	amount += animalCat(10)
	amount += animalMouse(5)
	amount += animalSpider(8)

	fmt.Println("amount of food to buy:", amount, "Kg")
}

func Dog(amountDogs float64) float64 {
	total := amountDogs * 10
	return total
}

func Cat(amountCats float64) float64 {
	total := amountCats * 5
	return total
}

func Mouse(amountMouse float64) float64 {
	total := (amountMouse / 1000) * 250
	return total
}

func Spider(amountSpider float64) float64 {
	total := (amountSpider / 1000) * 150
	return total
}

func animal(animal string) (func(float64) float64, string) {

	switch animal {
	case "dog":
		functionDog := Dog
		return functionDog, ""
	case "cat":
		functionCat := Cat
		return functionCat, ""
	case "mouse":
		functionMouse := Mouse
		return functionMouse, ""
	case "spider":
		functionSpider := Spider
		return functionSpider, ""
	default:
		return nil, "The animal don't exist"
	}
}
