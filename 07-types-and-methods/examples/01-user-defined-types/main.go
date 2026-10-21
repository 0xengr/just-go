package main

import "fmt"

// NOTE: Types are executable documentation: they name concepts and let the compiler
// reject mistakes that primitive types cannot express.

// NOTE: A defined type has its own identity, even when it shares an underlying
// type with another type.
type UserID int
type Score int
type Email string
type HighScore Score

// NOTE: A type alias is only another name for the same type.
type ID = int

// NOTE: iota is used for related constants of a defined type.
type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
)

func main() {
	var id UserID = 101
	var score Score = 95
	var email Email = "mohammad@example.com"

	fmt.Println(id, score, email)

	// NOTE: Distinct defined types are not interchangeable.
	// var x Score = id // COMPILE ERROR
	fmt.Println("score + 10:", score+10)

	// NOTE: Alias is the same type; defined type needs conversion.
	var n int = 10
	var alias ID = n // OK
	var converted UserID = UserID(n)
	fmt.Println("alias:", alias, "converted:", converted)

	// NOTE: The underlying type of HighScore is int, not Score.
	var hs HighScore = 200
	hs = HighScore(score)
	fmt.Println("high score:", hs)

	// NOTE: Untyped constants assign directly; typed values need conversion.
	var literal UserID = 10
	fmt.Println("literal:", literal, "category:", Personal)
}
