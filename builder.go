package main

import (
	"fmt"
)

type NotificationBuilder struct {
	title    string
	subtitle string
	message  string
	image    string
	icon     string
	priority int
	notype   string
}

func newnotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}

}

func (nb *NotificationBuilder) setTitle(title string) {

	nb.title = title

}

func (nb *NotificationBuilder) setsubtitle(subtitle string) {

	nb.subtitle = subtitle

}
func (nb *NotificationBuilder) setmessage(message string) {

	nb.message = message
}

func (nb *NotificationBuilder) setimage(image string) {

	nb.image = image

}

func (nb *NotificationBuilder) seticon(icon string) {

	nb.icon = icon

}

func (nb *NotificationBuilder) setpriority(pri int) {

	nb.priority = pri

}
func (nb *NotificationBuilder) setType(notype string) {

	nb.notype = notype

}

func (nb *NotificationBuilder) Build() (*NotificationBuilder, error) {

	if nb.icon != "" && nb.subtitle == "" {

		return nil, fmt.Errorf("You need specify a subtitle when using an icon ")
	}

	if nb.priority > 5 {
		return nil, fmt.Errorf("priority must be 0 to 5")
	}

	return &NotificationBuilder{

		title:    nb.title,
		subtitle: nb.subtitle,
		message:  nb.message,
		priority: nb.priority,
		image:    nb.image,
		icon:     nb.icon,
		notype:   nb.notype,
	}, nil

}
func main() {
	var bldr = newnotificationBuilder()

	bldr.setTitle("new Notification")
	bldr.seticon("incon.png")
	bldr.setsubtitle("This is a sub title")
	bldr.setimage("image.jpg")
	bldr.setpriority(5)
	bldr.setmessage("This is a basdic notication")
	bldr.setType("Alert")

	notif, error := bldr.Build()

	if error != nil {

		fmt.Println("Error creating the notification bar :", error)
	} else {

		fmt.Printf("Notification : %+v", notif)
	}

}
