package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Webrequests in Golang")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type: %T\n", response)
	fmt.Printf("%+v\n\n", *response)

	defer response.Body.Close() // Callers responsibility to close this connection

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n\n", string(databytes))

	statuscode := response.Status
	fmt.Printf("Stauscode: %v\n\n", statuscode)

	header := response.Header
	//format string Header into good looking json format
	headerJson, err := json.MarshalIndent(header, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(headerJson))
}
