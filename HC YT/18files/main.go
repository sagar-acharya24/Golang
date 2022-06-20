package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file - LearnCodeOnline.in"

	files, err := os.Create("./mylcogofile.txt") // os package creating .txt new text file, ./ means current directory

	//if err != nil{
	//	panic(err)
	//}
	checkNilErr(err)

	length, err := io.WriteString(files, content)

	checkNilErr(err)

	fmt.Println("length is : ", length)

	defer files.Close() //incase if we are going to write later

	readFile("./mylcogofile.txt")
}

//Read the .txt file which is created

func readFile(filename string){
	//reading this file it wont be in string it will be in byte format important

	databyte, err := ioutil.ReadFile(filename)

	checkNilErr(err)

	//Reading will happen in byte format so we have to convert it to string

	fmt.Println("Text data inside the file is \n", string(databyte)) 
}
	// Instead of calling the error nil function again and again 
	// we can create in function n call that where ever neccessary

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

