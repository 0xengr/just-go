package main

import (
	"errors"
	"fmt"
)

// NOTE: The purpose of this program is to calculate the pay for the week.

var (
	ErrHourlyRate   = errors.New("invalid hourly rate")
	ErrHourlyWorked = errors.New("invalid hours worked per week")
)

func payDay(hoursWorked, hourlyRate int) (int, error) {
	if hourlyRate < 10 || hourlyRate > 75 {
		return 0, ErrHourlyRate
	}
	if hoursWorked < 0 || hoursWorked > 80 {
		return 0, ErrHourlyWorked
	}

	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		regularPay := hoursWorked * hourlyRate
		return regularPay + overTime, nil
	}

	return hoursWorked + hourlyRate, nil
}

func main() {

	pay, err := payDay(81, 50)
	if err != nil {
		fmt.Println(err)
	}

	pay, err = payDay(80, 5)
	if err != nil {
		fmt.Println(err)
	}

	pay, err = payDay(80, 50)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pay)

}
