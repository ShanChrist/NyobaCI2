package main

import "fmt"

func main() {
	var x = 0
	if x = 1; x == 1 { // This line has a mistake, should use the := operator instead of = to declare and assign a new variable.
		fmt.Println("x is equal to 1")
	}
}
