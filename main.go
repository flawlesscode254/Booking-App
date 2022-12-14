package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	name := "Duncan"
	fmt.Println(name)

	var email = "duncanii414@gmail.com"
	fmt.Println(email)

	num := 4
	fmt.Println(num)

	var count = 7
	fmt.Println(count)

	var username int = 34
	fmt.Println(username)

	var state bool = false
	fmt.Println(state)

	//var identity string
	//
	//fmt.Println("What is your name:")
	//fmt.Scan(&identity)
	//
	//fmt.Printf("Your username is: %v", identity)

	// Sum calculations
	a := 12
	b := 13

	// Operations
	fmt.Println(a - b)
	fmt.Println(a + b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// Arrays implementation in Go
	var arrayStrings = [4]string{"A", "B", "C", "D"}
	fmt.Println(arrayStrings)

	var arrayInts = [4]int{1, 2, 3, 4}
	fmt.Println(arrayInts)

	var arrayBools = [2]bool{false, true}
	fmt.Println(arrayBools)

	var counterArr [3]int
	counterArr[2] = 4
	fmt.Println(counterArr)
	fmt.Println(len(counterArr))

	// No specified size array
	var trippy = []int{1, 2, 3, 4, 5}
	trippy = append(trippy, 23)
	fmt.Println(trippy)

	// Grouped variables
	var (
		f string = "A"
		g int    = 4
	)
	fmt.Println(f, g)

	// Go constants
	const attr = "Swelled"
	fmt.Println(attr)

	// Floats
	var floater float32 = 3.43
	fmt.Println(floater)

	// If statements
	if 20 < 45 {
		fmt.Println(true)
	}

	fmt.Println(20 < 45)
	fmt.Println(45 < 20)

	var sa = 20
	var se = 45

	fmt.Println(sa > se)

	if sa < se {
		fmt.Println(true)
	} else if sa > 100 {
		fmt.Println(false)
	} else {
		fmt.Println(true)
	}

	// Switch Statements
	var won = 23
	switch won {
	case 24:
		fmt.Println(true)
	}

	// For loops implementation
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// For loops from array
	var test = [5]int{1, 2, 3, 4, 5}
	for i := 1; i < len(test)+1; i++ {
		fmt.Println(i)
	}

	// Objects in GO
	type User struct {
		name  string
		email string
	}

	var person1 User
	person1.name = "Duncan"
	person1.email = "denzme414@gmail.com"

	fmt.Println(person1)
}
