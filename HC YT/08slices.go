package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList []string = []string{"Apple", "Tomamto", "Peaches"}

	fmt.Println(fruitList)

	fmt.Printf("Data type of fruitlist: %T\n", fruitList)

	fruitList = append(fruitList, "mango", "banana")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])

	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 222
	highScores[1] = 987
	highScores[2] = 456
	highScores[3] = 666
	//highScores[4] = 777

	//fmt.Println(highScores)

	highScores = append(highScores, 111, 333)

	//fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores)) //Boolean

	var courses []string = []string{"reactjs", "javascript", "swift", "python", "ruby"}

	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)
}
