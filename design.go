package main

import (
	"fmt"
)

type Dog interface {
	Bark()
	walk()
}

type Dogimpl struct {
}

func (d *Dogimpl) Bark() {

	fmt.Println("Woof")
}

func (d *Dogimpl) walk() {

	fmt.Println("Walk")

}

type Bulldog struct {
	Dog
}

type Flayable interface {
	Fly()
}

type Flayableimpls struct {
}

func (f *Flayableimpls) Fly() {

	fmt.Println("Flying")

}

type RobotDog struct {
	Dog
	Flayable
}

func main() {

	dog := new(Dogimpl)
	Bulldog := new(Bulldog)

	Bulldog.Dog = dog
	Bulldog.walk()
	Bulldog.Bark()

	fly := new(Flayableimpls)
	robotDog := new(RobotDog)

	robotDog.Dog = dog
	robotDog.Flayable = fly
	robotDog.Bark()
	robotDog.walk()
	robotDog.Fly()

	/*
		    Another way
		    var jj Dogimpl

			var yy Bulldog

			yy.Dog = &jj
			yy.Bark()
			yy.walk()

			var ff Flayableimpls
			var rr RobotDog
			rr.Dog = &jj
			rr.Flayable = &ff
			rr.Bark()
			rr.walk()
			rr.Fly()




	*/
}
