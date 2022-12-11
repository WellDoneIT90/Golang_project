package main

import "fmt"

func main() {
	fmt.Println("This is an example for maps")

	languages := make(map[string]string)

	languages["PY"] = "Python"
	languages["JV"] = "Java"
	languages["GO"] = "Golang"

	fmt.Println(languages)
	fmt.Printf("My favourite languages are: %s\n", languages)

	delete(languages, "JV")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("Fory key %v, value is %v\n", key, value)
	}
}
