package main

import (
	"fmt"
	"strconv"
)

func main() {
	exercise1()
	exercise2()
	exercise3()
}

// Read two numbers and print the sum
func exercise1() {
	var n1 int
	var n2 int

	fmt.Println("Type a X")
	fmt.Scanln(&n1)
	fmt.Println("Type a Y")
	fmt.Scanln(&n2)

	var sum = n1 + n2

	fmt.Println("X + Y = " + strconv.Itoa(sum))
}

// Read a number and print if it is greater than, less than or equal to 100
func exercise2() {
	var num int

	fmt.Println("Type a number")
	fmt.Scanln(&num)

	if num > 100 {
		fmt.Println("Greater than 100")
	} else if num < 100 {
		fmt.Println("Less than 100")
	} else {
		fmt.Println("Equal to 100")
	}
}

// Read a float 64 that represents a radius of a circle and calculate your area
// Consider PI = 3.14159
func exercise3() {
	const pi = 3.14159

	var r float64

	fmt.Println("Type a circle radius")
	fmt.Scan(&r)

	a := pi * (r * r)

	fmt.Printf("Area = %.4f\n", a)
}
