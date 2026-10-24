package main

import (
	"errors"
	"fmt"
)

// NOTE: Demonstrate a notification system without interfaces first,
// * then make it scalable with interfaces.

// NOTE: Building a notification system without interfaces.
func SendEmail(message string) {
	fmt.Println("Email:", message)
}
func SendSMS(message string) {
	fmt.Println("SMS:", message)
}
func SendPush(message string) {
	fmt.Println("Push:", message)
}

// NOTE: The caller has to know every channel, so this switch has to grow
// * every time a new channel is added.
func NotifyUser(channel, message string) {
	switch channel {
	case "Email":
		SendEmail(message)
	case "SMS":
		SendSMS(message)
	case "Push":
		SendPush(message)
	}
}

// NOTE: Building a notification system with interfaces.
// * Any type with a Send method satisfies Notifier, so NotifyTheUser
// * never has to change when a new channel is added.
type Notifier interface {
	Send(to, message string) error
}

type emailNotifier struct{}
type smsNotifier struct{}
type pushNotifier struct{}

func (emailNotifier) Send(to, message string) error {
	fmt.Printf("Sending EMAIL to %s: %s\n", to, message)
	return nil
}
func (smsNotifier) Send(to, message string) error {
	if to == "" {
		return errors.New("phone number is required")
	}
	fmt.Printf("Sending SMS to %s: %s\n", to, message)
	return nil
}
func (pushNotifier) Send(to, message string) error {
	fmt.Printf("Sending PUSH notification to device %s: %s\n", to, message)
	return nil
}

// NOTE: A function accepting an interface as a parameter.
func NotifyTheUser(n Notifier, recipient, message string) error {
	return n.Send(recipient, message)
}

func main() {
	NotifyUser("Email", "Hello from Email")
	NotifyUser("SMS", "Hello from SMS")
	NotifyUser("Push", "Hello from Push")
	NotifyUser("random", "Hello from random") // no case matches, so nothing is sent

	_ = NotifyTheUser(
		emailNotifier{},
		"0xengr@gmail.com",
		"Your order has shipped",
	)

	_ = NotifyTheUser(
		smsNotifier{},
		"+9779800000000",
		"Your verification code is 123456",
	)

	_ = NotifyTheUser(
		pushNotifier{},
		"device-42",
		"You have a new message",
	)
}
