package main

import "fmt"

func main() {
	fmt.Println("This is an example on slices")

	var fruitList = []string{"Apple", "Grape", "Peach"}
	fmt.Printf("Type of fruit list is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)
	fmt.Println(fruitList[1:3])

	// delete one element from slices with append

	index := 2

	fruitList = append(fruitList[:index], fruitList[index+1:]...)
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 777

	highScores = append(highScores, 777)

	fmt.Println(highScores)

}
