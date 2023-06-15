package main

import "fmt"

type grouptrip interface {
	setpackage(pack string)
	setplace(place string)
	setstay(duration string)
	getpackage() string
	getplace() string
	getstay() string
}

type trip struct {
	packagename string
	Visitplace  string
	stay        string
}

type user struct {
	Departurdate string
	Adult        int
	childeren    int
	Amount       int
	visit        grouptrip
}

func userdetails() *user {

	return &user{}
}

func (a *user) setdate(date string) {
	a.Departurdate = date
}

func (a *user) setpeople(person int) {
	a.Adult = person

	for i := 1; i <= person; i++ {

		fmt.Println("Enter name    :")
		fmt.Println("Enter DOB     :")
		fmt.Println("Enter Gender  :")
		fmt.Println("Enter MailID  :")
		fmt.Println()
	}

}
func (a *user) setchild(child int) {
	a.childeren = child
}

func (a *user) setvisit(confirmed grouptrip) {
	a.visit = confirmed
	//a.visit.getpackage()

}

func (a *user) details() (*user, error) {

	if a.Adult > 6 {

		return nil, fmt.Errorf("Adult must be up to 5")
	}

	return &user{

		Departurdate: a.Departurdate,
		Adult:        a.Adult,
		childeren:    a.childeren,
		Amount:       a.Adult * 5000,
		visit:        a.visit,
	}, nil

}

func (p *trip) setpackage(pack string) {

	p.packagename = pack
}

func (p *trip) setplace(place string) {

	p.Visitplace = place
}

func (p *trip) setstay(duration string) {

	p.stay = duration
}

func (p *trip) getpackage() string {

	return p.packagename
}

func (p *trip) getplace() string {

	return p.Visitplace
}

func (p *trip) getstay() string {

	return p.stay
}

type familytrip struct {
	trip
}

//func (a *familytrip) call() {

//fmt.Sprintf("Trip name si %s", a.packagename)
//}

func maketrip(nameofthetrip string, visit string, staying string) grouptrip {

	return &familytrip{

		trip: trip{
			packagename: nameofthetrip,
			Visitplace:  visit,
			stay:        staying,
		},
	}

}

func executethetrip(packname string, area string, duration string) (grouptrip, error) {

	return maketrip(packname, area, duration), nil
}

func tripdetails(gt grouptrip) {

	fmt.Printf("_________________________________________________\n")
	fmt.Printf("%s\n", gt)
	fmt.Printf("Package  Name  :%s\n", gt.getpackage())
	fmt.Printf("Visiting Place :%s\n", gt.getplace())
	fmt.Printf("stay           :%s\n", gt.getstay())
	fmt.Printf("________________________________________________\n")

}

func payment(pay chan int, click string) {

	if click == "clicK" {

		fmt.Println("Total :", <-pay)
	} else {
		fmt.Println("cancelled")
	}
	//fmt.Println(click)
}
func main() {

	visit1, _ := executethetrip("Dubai Group Departure ex Chennai", "Sharjah | Abu Dhabi | Dubai", "4 Nights / 5 Days")
	visit2, _ := executethetrip("Magical Malaysia Group Tour ex Chennai", " Malaysia", "3 Nights / 4 Days")
	visit3, _ := executethetrip("Europe Winter Group Departur ex Chennai", "Zurich | Axams | Padova | Rome | Arezzo", "10 Nights / 11 Days")

	fmt.Println("Visit1", visit1)

	tripdetails(visit1)
	tripdetails(visit2)
	tripdetails(visit3)

	var vi = userdetails()
	vi.setdate("11-jan-2034")
	vi.setpeople(3)
	vi.setchild(0)
	vi.setvisit(visit1)

	notify, error := vi.details()

	if error != nil {

		fmt.Printf("An Error is : %s", error)
	} else {

		//fmt.Printf("Viewcart :%+v", notify, notify.visit.getpackage())
		fmt.Println("Viewcart :", vi, notify.Amount, notify.visit.getpackage())
	}

	var pay chan int = make(chan int)
	go payment(pay, "clicK")
	pay <- notify.Amount

}
