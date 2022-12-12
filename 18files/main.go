package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in Golang")

	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checlNilErr(err)

	length, err := io.WriteString(file, content)
	checlNilErr(err)

	fmt.Println("Length is:", length)
	defer file.Close()

	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	// in go 1.16 it was used as ioutil.ReadFile(filename string) ([]byte, error),
	// from go 1.19 we will use new func ReadFile(filename string) ([]byte, error) from os package
	databyte, err := os.ReadFile(filename)
	checlNilErr(err)

	fmt.Println("Text data inside the file is: \n", string(databyte))

}

func checlNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
