package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Print(a.food)
}

func (a Animal) Move() {
	fmt.Print(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Print(a.noise)
}

func (a Animal) HandleRequest(req, name string ) {
	switch req {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Print("Invalid request! You may try:")
		fmt.Printf("\n>%s eat", name)
	}
}

func main() {
	animals := make(map[string]Animal)

	animals["cow"] = Animal{
		"grass",
		"walk",
		"moo",
	}

	animals["bird"] = Animal{
		"worms",
		"fly",
		"peep",
	}

	animals["snake"] = Animal{
		"mice",
		"slither",
		"hsss",
	}

	var inp, req string
	var animal Animal
	var exists bool

	fmt.Println("-------------------------- WELCOME -------------------------")
	fmt.Println("Input your command as two word or type 'q' for exit. i.e.")
	fmt.Println(">snake eat")
	fmt.Println(">cow speak")
	fmt.Println(">bird move")
	fmt.Println(">q")
	fmt.Println("============================================================")

	for {
		fmt.Print(">")
		fmt.Scanf("%s %s", &inp, &req)

		inp = strings.ToLower(inp)
		if inp == "q" {
			fmt.Println("Thank you!")
			break
		}

		if animal, exists = animals[inp]; exists {
			animal.HandleRequest(req, inp)
			fmt.Println("")
		} else {
			fmt.Println("Invalid animal! Try again")
		}
	}
}
