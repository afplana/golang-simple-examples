package main

import (
	"fmt"
)

// GenDisplaceFn clalculate displacement
func GenDisplaceFn(a, v, u float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (0.5*a*t*t + v*t + u)
	}
	return fn
}

func main() {
	var a, u0, s0, t float64
	fmt.Println("Enter acceleration, initial velocity and initial displacement: ")
	fmt.Scan(&a)
	fmt.Scan(&u0)
	fmt.Scan(&s0)
	getDisplacement := GenDisplaceFn(a, u0, s0)
	fmt.Println("Enter time: ")
	fmt.Scan(&t)
	fmt.Println(getDisplacement(t))
}
