package main

import "fmt"

type Data struct {
	name   string
	age    string
	gender string
}

func (p *Data) username(na string) string {

	p.name = na

	return p.name
}

func (p *Data) userage(ag string) string {

	p.age = ag

	return p.age
}
func (p *Data) usergender(gen string) string {

	p.gender = gen

	return p.gender
}

func (p *Data) getting() (*Data, error) {

	return &Data{

		name:   p.name,
		age:    p.age,
		gender: p.gender,
	}, nil
}

func collecting() *Data {

	return &Data{}
}

func main() {
	v := collecting()

	v.username("Vijay")
	v.userage("24")
	v.usergender("Male")
	vi, err := v.getting()

	if err != nil {

		fmt.Println("Something went wrong :", err)
	} else {

		fmt.Printf("User data is %+v:", vi)
	}
}
