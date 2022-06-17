package main

import (
	"fmt"
	"sync"
	"time"
)

func eyes(w *sync.WaitGroup) {
	fmt.Println("eyes")
	w.Done()
}

func ears(w *sync.WaitGroup) {
	fmt.Println("ears")
	w.Done()
}

func talk(w *sync.WaitGroup) {
	fmt.Println("talk")
	w.Done()
}

func main() {
	fmt.Println("main function")

	var wg sync.WaitGroup

	wg.Add(1)

	go eyes(&wg)
	
	wg.Add(1)

	go ears(&wg)

	wg.Add(1)

	go talk(&wg)

	fmt.Println("before wait")

	wg.Wait()

	time.Sleep(1)

}
