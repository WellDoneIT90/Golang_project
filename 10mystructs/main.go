package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	//no inheritance in golang; no super or parent

	welldone := User{"Welldone", "welldone@welldone.it", true, 32}
	fmt.Println(welldone)
	fmt.Printf("%+v\n", welldone)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
