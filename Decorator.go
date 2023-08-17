package main

import (
	"fmt"
)

type showroom interface {
	sale()
}
type Modelcars struct{}

func (car *Modelcars) sale() {

	fmt.Println("We are always great")

}

type Cars struct {
	secondhand *Modelcars
}

func (car *Cars) sale() {

	if car.secondhand == nil {

		fmt.Println("We don't have the data")
	} else {

		fmt.Println("yes we have ")
		car.secondhand.sale()

	}

}

func main() {

	// This is one way

	var vj = &Modelcars{}

	var Rj = &Cars{}

	Rj.sale()

	Rj.secondhand = vj

	Rj.sale()

	/*

		                This is one way
				//var vj Modelcars

				var rj Cars

				rj.sale()

				rj.secondhand = &Modelcars{}

				rj.sale()

				fmt.Println(rj)
				
				The ProcessDecorator struct is a wrapper around the ProcessClass struct. It adds new functionality to the
				 ProcessClass struct without having to modify the original code.

	*/

}
