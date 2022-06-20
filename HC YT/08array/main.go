package main

import "fmt"

func main() {
	fmt.Println("Welcome to array")

	var fruitList [4]string

	fruitList[0] = "Apple"

	fruitList[1] = "orange"

	fruitList[3] = "peach"

	fmt.Println("List of fruits: ", fruitList)

	fmt.Println("length of List of fruits: ", len(fruitList))

	var veggy [3]string = [3]string{"cabbage","beans","mushroom"}

	fmt.Println("list of veg:", veggy)

	fmt.Println("len of list of veg:", len(veggy))

}
