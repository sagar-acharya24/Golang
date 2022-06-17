package main

import (
	"fmt"
)

func main() {
	fmt.Println("Data structure1")

	var godphoto [4]string = [4]string{"ganapa", "lakshmi", "shiva", "hari"}

	fmt.Println(godphoto[1])

	var humanwishes []string = []string{"own jet", "house", "luxury car", "bike", "money", "foreign trip"}

	humanwishes= append(humanwishes, "sky diving")

	fmt.Println(humanwishes[:])

	var yash map[string]string =make(map[string]string,10)

	yash["teen"]="smart"

	yash["young"]= "super smart"

	yash["present"]="hero"

	yash["old"]="grandpa"

	fmt.Println(yash)
}
