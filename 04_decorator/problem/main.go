package main

import "fmt"

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

// TODO: If I want more notifiers, I have to define more types to combine them!!!
type EmailSMSNotifier struct {
	emailNotifier EmailNotifier
	smsNotifier   SMSNotifier
}

func (notifier EmailSMSNotifier) Send(message string) {
	notifier.emailNotifier.Send(message)
	notifier.smsNotifier.Send(message)
}

type Service struct {
	notifier Notifier
}

func (s Service) SendNotification(mesage string) {
	s.notifier.Send(mesage)
}

func main() {
	s := Service{
		notifier: EmailSMSNotifier{},
	}

	s.SendNotification("Hello World")
}
