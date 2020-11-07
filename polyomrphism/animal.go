package main

import (
	"fmt"
)

// Animal contains basic info about animals
type Animal struct {
	food, locomotion, noise string
}

// Eat return type of food based on the animal
func (a *Animal) Eat() {
	fmt.Println(a.food)
}

// Move return type of movement based on the animal
func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

// Speak return type of sound based on the animal
func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func (a *Animal) initAnimal(f, l, n string) {
	a.food = f
	a.locomotion = l
	a.noise = n
}

func animalInfo(a Animal, info string) {
	switch info {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Println("Invalid information request!. Provide a valid info request")
	}
}

func main() {
	var reqAnimal, reqInfo string
	var cow, bird, snake Animal

	cow.initAnimal("grass", "walk", "moo")
	bird.initAnimal("worms", "fly", "peep")
	snake.initAnimal("mice", "slither", "hsss")

	for {
		fmt.Println("[*]Please sumbit a single request in the format: animal information")
		fmt.Println("[*]Available animals: [cow, bird, snake] | Avaliable Info: [eat, move, speak]")
		fmt.Print(">")
		fmt.Scanf("%s %s", &reqAnimal, &reqInfo)
		fmt.Print("[*]Information response: ")

		switch reqAnimal {
		case "cow":
			animalInfo(cow, reqInfo)
		case "bird":
			animalInfo(bird, reqInfo)
		case "snake":
			animalInfo(snake, reqInfo)
		default:
			fmt.Println("Invalid animal type in request!. Please provide a valid animal name")
		}

		fmt.Println()
	}
}
