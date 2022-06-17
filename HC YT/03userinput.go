package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "welcometo user input"

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter rating for our pizza: ")

	input , _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	fmt.Printf("Type of rating : %T", input)
}
