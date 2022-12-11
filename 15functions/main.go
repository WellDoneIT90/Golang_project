package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions in Golang")
	companyName()
	//company name with string
	companyName_string("WelldoneIT")
	fmt.Println("The volume is: ", volume(5, 5, 5))

	//dynamic value function call
	total, message := proadder(5, 4, 7, 3, 6, 100, -50)
	fmt.Printf("The added value is: %v and message: %v\n", total, message)
}

func volume(x int, y int, z int) int {
	return x * y * z
}

func proadder(values ...int) (int, string) {
	total := 0

	fmt.Println(values)

	for _, x := range values {
		total += x
	}

	return total, "Here could be your commercial!"
}

func companyName() {
	fmt.Println("My company name is WelldoneIT")
}

func companyName_string(name string) {
	fmt.Println("My company name is", name)
}
