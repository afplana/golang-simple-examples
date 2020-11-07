package main

import (
	"encoding/json"
	"fmt"
)

// Person can be used by everybody
type Person struct {
	name  string
	addr  string
	phone string
}

// P is a point in the plane representation
type P struct {
	x int
	y int
}

var f func(x int) int

func test(x int) int {
	return x + 1
}

func fA() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func modif(arr []int) {
	arr[0] = arr[0] + 1
}

func main() {
	p := P{x: 45, y: 67}
	var p1 *P
	value, _ := json.Marshal(p)
	fmt.Println(p)
	fmt.Println(value)
	e := json.Unmarshal(value, &p1)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(p1)
	f = test
	m := f(5)

	fmt.Println(m)
	x := []int{1, 2, 3}
	modif(x)
	fmt.Println(x)

	fB := fA()
	fmt.Println(fB())
	fmt.Println(fB())

	i := 1

	fmt.Print(i)

	i++

	defer fmt.Print(i + 1)

	fmt.Print(i)
}
