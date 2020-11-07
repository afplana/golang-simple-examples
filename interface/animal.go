package main

import "fmt"

// Animal defines actions
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow define an animal type
type Cow struct {
	name, food, locomotion, noise string
}

// Bird define an animal type
type Bird struct {
	name, food, locomotion, noise string
}

// Snake define an animal type
type Snake struct {
	name, food, locomotion, noise string
}

// Eat of the cow
func (c *Cow) Eat() {
	fmt.Println(c.food)
}

// Move of the cow
func (c *Cow) Move() {
	fmt.Println(c.locomotion)
}

// Speak of the cow
func (c *Cow) Speak() {
	fmt.Println(c.noise)
}

// Eat of the Bird
func (b *Bird) Eat() {
	fmt.Println(b.food)
}

// Move of the Bird
func (b *Bird) Move() {
	fmt.Println(b.locomotion)
}

// Speak of the Bird
func (b *Bird) Speak() {
	fmt.Println(b.noise)
}

// Eat of the Snake
func (s *Snake) Eat() {
	fmt.Println(s.food)
}

// Move of the Snake
func (s *Snake) Move() {
	fmt.Println(s.locomotion)
}

// Speak of the Snake
func (s *Snake) Speak() {
	fmt.Println(s.noise)
}

func printStartInfo() {
	fmt.Println("[*]Sumbit a single request of type newanimal or query. Type q for exit!")
	fmt.Println("[!]New Animal request structure: newanimal arbitrary_name animal_type")
	fmt.Println("[!]Query request structure: query animal_name animal_info")
	fmt.Println("[!]Animal Type: [cow, bird, snake] | Avaliable Info: [eat, move, speak]")
}

func buildAnimal(atype string, name string) Animal {
	switch atype {
	case "cow":
		return &Cow{name, "grass", "walk", "moo"}
	case "bird":
		return &Bird{name, "worms", "fly", "peep"}
	case "snake":
		return &Snake{name, "mice", "slither", "hsss"}
	default:
		fmt.Println("Invalid animal type in request!. Please provide a valid animal name")
	}

	return nil
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
	var rqtype, name, info string
	animals := make(map[string]Animal)

	printStartInfo()

	for {
		fmt.Print(">")
		fmt.Scanf("%s %s %s", &rqtype, &name, &info)

		switch rqtype {
		case "newanimal":
			animals[name] = buildAnimal(info, name)
			fmt.Println("Created it!")
		case "query":
			if animals[name] != nil {
				animalInfo(animals[name], info)
			} else {
				fmt.Println("No animal under name " + name + " registred!")
			}
		case "q":
			fmt.Println("Thank you!")
			return
		default:
			fmt.Println("Invalid request! Command should start with newanimal || query")
		}

		fmt.Println()
	}
}
