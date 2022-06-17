package main

import (
	"fmt"
	"time"
)

func ridingvehicle(r chan bool) {
	fmt.Println("riding vehicle", <-r) //passing the value from channel to reference
}

func eyes(r chan bool) {
	fmt.Println("eyes", <-r)
}

func headset(r chan bool) {
	fmt.Println("head", <-r)
}

func main() {
	fmt.Println("Hello world!")

	road := make(chan bool,3)

	go ridingvehicle(road)

	fmt.Println("write into resource")

	road <- true //saying road available true

	go eyes(road)

	road <- true

	go headset(road)

	road <- true

	time.Sleep(1)

}
