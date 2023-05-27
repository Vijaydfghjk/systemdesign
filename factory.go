package main

import "fmt"

type ipublication interface {
	setname(name string)
	setpages(pages int)
	setpublisher(publish string)
	getname() string
	getpage() int
	getpubshier() string
}

type publication struct {
	name      string
	pages     int
	publisher string
}

func (p *publication) setname(name string) {

	p.name = name
}

func (p *publication) setpages(pages int) {
	p.pages = pages
}

func (p *publication) setpublisher(publish string) {

	p.publisher = publish
}

func (p *publication) getname() string {

	return p.name
}

func (p *publication) getpage() int {

	return p.pages
}

func (p *publication) getpubshier() string {

	return p.publisher
}

type newspaper struct {
	publication
}

func (n newspaper) string() string {

	return fmt.Sprint("This is newpaper named %s", n.name)
}

func createnewpaper(name string, pages int, publisher string) ipublication {

	return &newspaper{

		publication: publication{

			name:      name,
			pages:     pages,
			publisher: publisher,
		},
	}
}

type magazine struct {
	publication
}

func (m *magazine) string() string {

	return fmt.Sprintf("Thi is a magazine named %s", m.name)
}

func createmagazine(name string, pages int, publisher string) ipublication {

	return &magazine{

		publication: publication{

			name:      name,
			pages:     pages,
			publisher: publisher,
		},
	}
}

func newpublication(pubType string, name string, pg int, pub string) (ipublication, error) {

	if pubType == "newspaper" {

		return createnewpaper(name, pg, pub), nil
	}
	if pubType == "magazine" {

		return createmagazine(name, pg, pub), nil
	}
	return nil, fmt.Errorf("No such publication type")
}

func main() {

	mag1, _ := newpublication("magazine", "tyme", 50, "The Tymes")
	mag2, _ := newpublication("magazine", "Lyfe", 40, "The Lyfs")
	mag3, _ := newpublication("newspaper", "The Hindu", 55, "The indu")
	mag4, _ := newpublication("newspaper", "India Today", 60, "India Today")

	pubdetails(mag1)
	pubdetails(mag2)
	pubdetails(mag3)
	pubdetails(mag4)

}

func pubdetails(pub ipublication) {

	fmt.Printf("-------------------------------------------\n")
	fmt.Printf("%s\n", pub)
	fmt.Printf("Name :%s\n", pub.getname())
	fmt.Printf("Pages :%d\n", pub.getpage())
	fmt.Printf("Publisher :%s\n", pub.getpubshier())
	fmt.Printf("----------------------------------------\n")

}
