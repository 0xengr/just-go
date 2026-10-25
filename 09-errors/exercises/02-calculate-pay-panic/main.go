package main

import (
	"errors"
	"fmt"
)

// NOTE: This program is similar to ./01-calculate-pay, but instead of
// * returning an error, it panics. Hours over 40 are paid at twice the rate.

var (
	ErrHourlyRate  = errors.New("invalid hourly rate")
	ErrHoursWorked = errors.New("invalid hours worked per week")
)

func payDay(hoursWorked, hourlyRate int) int {
	report := func() {
		fmt.Printf("HourlyRate: %d\nHoursWorked: %d\n", hourlyRate, hoursWorked)
	}
	defer report()

	if hourlyRate < 10 || hourlyRate > 75 {
		panic(ErrHourlyRate)
	}
	if hoursWorked < 0 || hoursWorked > 80 {
		panic(ErrHoursWorked)
	}

	if hoursWorked > 40 {
		regularPay := 40 * hourlyRate
		overTime := (hoursWorked - 40) * hourlyRate * 2
		return regularPay + overTime
	}

	return hoursWorked * hourlyRate

}

func main() {

	result := payDay(70, 10)
	fmt.Println(result)
}
