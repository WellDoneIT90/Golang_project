package main

import "fmt"

func main() {
	fmt.Println("This is a pointer example")

	// var ptr *int
	// fmt.Println("Value of pointer is", ptr)

	myNumber := 3

	var ptr = &myNumber

	*ptr = *ptr + 2

	fmt.Println("My number is: ", ptr)
	fmt.Println("My number is: ", *ptr)

}
