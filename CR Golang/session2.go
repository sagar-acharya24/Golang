package main

import (
	"fmt"
)

func movie() {
	var hero string = "sai"
	var villian string = "sasi"
	var budget string = "2cr"
	fmt.Println(hero, villian, budget)
}

type IGossip struct {
	name    string
	earning string
	house   string
}

type FGossip struct {
	name    string
	earning string
}

func main() {

	movie()

	var yash IGossip = IGossip{"venkat", "5cr", "duplex"}

	var laksh IGossip = IGossip{"venkat", "10cr", "duplex"}

	fmt.Println(yash, laksh)

	var navya FGossip = FGossip{"raj", "50000"}

	fmt.Println(navya)

	var radha sis = sis{bloodrelation{"rama", "seetha"}, "krish"}
	var pawan bro = bro{bloodrelation{"rama", "seetha"}, "lax"}
	fmt.Println(radha, pawan)
} 

type bloodrelation struct {
	father string
	mother string
}

type sis struct {
	bloodrelation
	husband string
}

type bro struct {
	bloodrelation
	wife string
}

