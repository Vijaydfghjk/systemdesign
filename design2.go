package main

import (
	"fmt"
)

type processer interface {
	process()
}

type hardware struct {
	control software
}

func (adp hardware) process() {

	fmt.Println("Good Processer")
	adp.control.convert()
}

type software struct {
	version int
}

func (add software) convert() {

	add.version = 5678

	fmt.Println("version is :", add)
}
func main() {

	var vj hardware
	var vi processer = vj
	vi.process()

	/*
		another way of accessing
			var Vijay processer = hardware{}
			Vijay.process()

	*/
}
