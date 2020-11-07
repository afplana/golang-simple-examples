package main

import (
	"fmt"
	"log"
)

// GenDisplaceFn Compute displacement as a function of time
func GenDisplaceFn(a, v, s float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*(t*t) + v*t + s
	}
}

func main() {
	var a, v, s, t float64

	fmt.Print("[*] Enter(float) acceletarion> ")
	if _, err := fmt.Scanf("%f", &a); err != nil {
		log.Fatal(err)
	}
	fmt.Print("[*] Enter(float) velocity> ")
	if _, err := fmt.Scanf("%f", &v); err != nil {
		log.Fatal(err)
	}
	fmt.Print("[*] Enter(float) initial displacement> ")
	if _, err := fmt.Scanf("%f", &s); err != nil {
		log.Fatal(err)
	}

	fn := GenDisplaceFn(a, v, s)

	fmt.Print("[*] Enter(float) Time> ")
	if _, err := fmt.Scanf("%f", &t); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Displacement: ", fn(t))
}
