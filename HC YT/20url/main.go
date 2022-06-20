package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to handling URL's in golang")

	fmt.Println(myurl)

	result, _ := url.Parse(myurl)    //url.Parse to extract info from URL (net/url package)

	fmt.Println(result.Scheme)

	fmt.Println(result.Host)

	fmt.Println(result.Path)

	fmt.Println(result.Port())

	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("The type of query param is: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams{		// for index,value := (syntax) , if we dont use range it won exe and it will show in problems
		fmt.Println("Param is: ", val)
	}

	//Creating URL important

	partsOfUrl := &url.URL{     //pass the address by using & url.URL{}
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hithesh",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)   //Open the in output by clicking ctrl +
}
