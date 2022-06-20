package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.jvtechnologies.co.in/"

func main() {
	fmt.Println("JVT web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()

	databyte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databyte)

	fmt.Println(content)

}
