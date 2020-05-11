package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a) // Initial

	var b, c int = 1, 2
	fmt.Println(b, c) // 1 2

	var d = true
	fmt.Println(d) // true

	var e int
	fmt.Println(e) // 0

	f := "apple"
	fmt.Println(f) // apple

	// Contants must have a declaration
	const g = "constant declaration"

	// This will come as error
	// const g string
	fmt.Println(g) // constant declaration
}
