package main

import "fmt"

func main() {

	DateCalculator()

	firstDate, secondDate := inputDates()

	message, _ := calculateDifference(firstDate, secondDate)

	fmt.Println(display(message))

}
