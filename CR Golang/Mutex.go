//Mutex

package main

import (
	"fmt"
	"sync"
	"time"
)

var tape string = "water"

func person(p *sync.Mutex) {
	p.Lock()

	tape = "drinking water"

	fmt.Println("Drinking water")

	p.Unlock()

}

func main() {
	fmt.Println("Mutex World")

	var sagar sync.Mutex

	go person(&sagar)

	time.Sleep(1)
}
