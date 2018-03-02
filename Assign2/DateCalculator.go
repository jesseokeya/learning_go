package main

import (
	"fmt"
	"math"
	"strconv"
)

// Calculator struct
type Calculator struct {
	firstDate  Date
	secondDate Date
	message    string
}

// DateCalculator is good
func DateCalculator() {
	fmt.Println("Date Calculator - it's all relative ")
	fmt.Println()
}

func inputDates() (Date, Date) {
	fmt.Println("Enter first date ")
	year := inputYear()
	month := inputMonth()
	day := inputDay()
	d1 := Date{
		day,
		month,
		year,
	}

	fmt.Println("Enter second date ")
	secondYear := inputYear()
	secondMonth := inputMonth()
	secondDay := inputDay()
	d2 := Date{
		secondDay,
		secondMonth,
		secondYear,
	}
	return d1, d2
}

func inputDate(d Date) Date {
	return d
}

func calculateDifference(d1 Date, d2 Date) (string, int) {
	firstDate := calcDays(d1)
	secondDate := calcDays(d2)
	difference := int(math.Abs(float64(firstDate - secondDate)))
	message := generateDisplay(d1, d2, difference)
	return message, difference
}

func display(s string) string {
	return s
}

// generateDisplay is a helper method to get displayMessage
func generateDisplay(d1 Date, d2 Date, difference int) string {
	message := ""

	if difference == 0 {
		message = displayDate(d1) + " is same as " + displayDate(d2)
	}

	if difference == 1 {
		if calcDays(d1) > calcDays(d2) {
			message = displayDate(d2) + " is " + strconv.Itoa(difference) + " day earlier than " + displayDate(d1)
		} else {
			message = displayDate(d1) + " is " + strconv.Itoa(difference) + " day earlier than " + displayDate(d2)
		}
	}

	if difference > 1 && difference < 7 {
		if calcDays(d1) > calcDays(d2) {
			message = displayDate(d2) + " is " + strconv.Itoa(difference) + " days earlier than " + displayDate(d1)
		} else {
			message = displayDate(d1) + " is " + strconv.Itoa(difference) + " days earlier than " + displayDate(d2)
		}
	}

	if difference > 7 && difference < 30 {
		difference = difference / 7
		if calcDays(d1) > calcDays(d2) {
			message = displayDate(d2) + " is " + strconv.Itoa(difference) + " weeks earlier than " + displayDate(d1)
		} else {
			message = displayDate(d1) + " is " + strconv.Itoa(difference) + " weeks earlier than " + displayDate(d2)
		}
	}

	if difference > 30 {
		difference = difference / 30
		if calcDays(d1) > calcDays(d2) {
			message = displayDate(d1) + " is " + strconv.Itoa(difference) + " months later than " + displayDate(d2)
		} else {
			message = displayDate(d1) + " is " + strconv.Itoa(difference) + " months later than " + displayDate(d2)
		}
	}

	return message
}
