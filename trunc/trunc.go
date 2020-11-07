package main

import "fmt"

func main() {
	var num float32
	fmt.Println("Float Truncator")
	fmt.Println("======================")
	fmt.Scan(&num)
	fmt.Printf("Truncated value: %.0f\n", num)
}
