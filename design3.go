package main

import (
	"fmt"
)

type continent interface {
	contries(a [5]float32, b [5]float32)
}

type area struct{}

func (p area) contries(a [5]float32, b [5]float32) {

	fmt.Println("Fix place")

}

type grouptowers interface {
	trip(x [5]float32, y [5]float32)
	resizeByFactor(factor int)
}

type group struct {
	x      [5]float32
	y      [5]float32
	place  area
	factor int
}

func (p group) trip(x [5]float32, y [5]float32) {

	fmt.Println("All places ")
	p.place.contries(p.x, p.y)
}

func (a group) resizeByFactor(factor int) {

	a.factor = factor
}

func main() {

	var a = [5]float32{1, 2, 3, 4, 5}
	var b = [5]float32{6, 7, 8, 9, 10}

	var vi grouptowers = group{a, b, area{}, 2}
	vi.trip(a, b)
	vi.resizeByFactor(2)

	/*
		   Another way of accessing memory

			var x = [5]float32{1, 2, 3, 4, 5}
			var y = [5]float32{5, 4, 3, 2, 1}
			var vj group

			var vi grouptowers = vj
			vi.trip(x, y)
			vi.vi.resizeByFactor(2)

	*/
}
