package main

import (
	"fmt"
	"sync"
	"time"
)

func ears() {
	fmt.Println("ears")
}

func eyes() {
	fmt.Println("eyes")
}

func main() {

	fmt.Println("MAIN FUNC")

	var wg sync.WaitGroup

	wg.Add(1);

	go ears()

	go eyes()

	wg.Wait();

	time.Sleep(1)

}
