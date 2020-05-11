package main

import (
	"fmt"
	"strconv"
)

func main() {
	exercise1()
	exercise2()
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
