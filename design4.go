package main

import (
	"fmt"
)

type icecream struct {
	Butterscotch int
}

func (p *icecream) ice(n string) {

	fmt.Println(p.Butterscotch, ":", n)
}

func (p *icecream) cream(price int) {

	p.Butterscotch = price

}

var own *icecream

func shop() *icecream {

	if own == nil {

		fmt.Println("Preparing icecream")
		return &icecream{}
	}
	fmt.Println("order preparing")
	return own
}

func main() {

	person := shop()

	person.cream(1)
	person.ice("I want ice cream")

	person = shop()
	person.cream(2)
	person.ice("I want ice cream cup")

	person = shop()
	person.cream(3)
	person.ice("give 4 ice cream")

}
