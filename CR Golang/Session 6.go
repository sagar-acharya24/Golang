//Abstraction - method (Using pointer)
package main

import (
	"fmt"
)

type Gossip struct {
	name    string
	earning string
	house   string

}

func (p *Gossip) news() {
	p.name = "venkat"
	p.earning = "5cr"
	p.house = "2cr"
}

func main() {
	var raj Gossip
	raj.news()
	fmt.Println(raj)
}
