package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	//no inheritance in golang; no super or parent

	welldone := User{"Welldone", "welldone@welldone.it", true, 32}
	fmt.Println(welldone)
	// fmt.Printf("%+v\n", welldone)
	welldone.GetStatus()
	welldone.SetEmail("welldoneit2@welldone.it")
	fmt.Println(welldone.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User active:", u.Status)
}

func (u User) SetEmail(email string) {
	//this changes the value only in the copy of User
	u.Email = email
	fmt.Println("Your new email adress is:", u.Email)
}
