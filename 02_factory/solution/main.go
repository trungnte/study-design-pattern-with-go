package main

import (
	"fmt"
)

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (EmailNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: Email)\n", message)
}

type SMSNotifier struct{}

func (SMSNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: SMS)\n", message)
}

type NotificationService struct {
	notifier Notifier
}

func (s NotificationService) SendNotification(message string) {
	s.notifier.Send(message)
}

func CreateNotifier(t string) Notifier {
	if t == "sms" {
		return SMSNotifier{}
	}
	return EmailNotifier{}
}

func main() {
	//! Problem: developer create new their struct and inject to our SDK
	//! Our SDK dont support their struct -> unknown behavior
	// s := NotificationService{
	// 	notifier: EmailNotifier{},
	// }

	//! Solution: We dont allow developer create new struct.
	//! We force them create Service with type
	// => Factory method
	s := NotificationService{
		notifier: CreateNotifier("sms"),
	}

	s.SendNotification("Hello World") // expected: Sending message: Hello World (Sender: SMS)

	s2 := NotificationService{
		notifier: CreateNotifier("abcxyz"),
	}

	//! only SMS and default email support if user pass other type!
	s2.SendNotification("Hello World") // expected: Sending message: Hello World (Sender: Email)
}
