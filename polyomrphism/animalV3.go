package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	sound      string
}

var animals map[string]Animal = map[string]Animal{
	"cow":   Animal{"grass", "walk", "moo"},
	"bird":  Animal{"worms", "fly", "peep"},
	"snake": Animal{"mice", "slither", "hsss"},
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.sound)
}

func main() {
	fmt.Printf("Enter animal and needed information\n")
	for {
		fmt.Print("> ")
		in := bufio.NewReader(os.Stdin)
		command, _ := in.ReadString('\n')
		animal := animals[strings.Split(command, " ")[0]]
		info := strings.Split(command, " ")[1]
		switch strings.TrimSpace(info) {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		}
	}
}
