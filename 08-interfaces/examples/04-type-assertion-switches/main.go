package main

import (
	"errors"
	"fmt"
)

// NOTE: Define the event payloads.

type UserCreated struct {
	UserID int
	Email  string
}
type PaymentFailed struct {
	UserID int
	Reason string
}
type OrderShipped struct {
	OrderID string
	Email   string
}

// NOTE: Process the events using a type switch.
// * `event.(type)` is only allowed inside a switch. In each case, `e` has the
// * case's concrete type, so `e.UserID`, `e.Email`, and so on are available
// * without a separate assertion.
// * The default case returns an error for any type this function does not know,
// * and `%T` reports that dynamic type instead of failing silently.
func processEvent(event any) error {
	switch e := event.(type) {
	case UserCreated:
		if e.UserID <= 0 || e.Email == "" {
			return errors.New("invalid user.created event")
		}
		fmt.Println("welcome email to:", e.Email)

	case PaymentFailed:
		if e.UserID <= 0 || e.Reason == "" {
			return errors.New("invalid payment-failed event")
		}

		fmt.Printf("payment failed for user %d: %s\n", e.UserID, e.Reason)

	case OrderShipped:
		if e.OrderID == "" || e.Email == "" {
			return errors.New("invalid order-shipped event")
		}
		fmt.Printf("shipping updates for order %s to %s\n", e.OrderID, e.Email)

	default:
		return fmt.Errorf("unsupported event type: %T", event)
	}

	return nil

}
func main() {
	// NOTE: One slice holds events of different concrete types; processEvent
	// * recovers each type with the switch above.
	events := []any{
		UserCreated{
			UserID: 42,
			Email:  "user@example.com",
		},
		PaymentFailed{
			UserID: 42,
			Reason: "card declined",
		},
		OrderShipped{
			OrderID: "ORD-1001",
			Email:   "user@example.com",
		},
	}

	for _, event := range events {
		if err := processEvent(event); err != nil {
			fmt.Println("Event processing error:", err)
		}
	}
}
