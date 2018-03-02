package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Date is exported
type Date struct {
	day   int
	month int
	year  int
}

// OurDate return *date
func OurDate(day int, month int, year int) Date {
	values := Date{
		day,
		month,
		year,
	}
	return values
}

func inputDay() int {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a day: ")
	day, err := scanner.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	result, newErr := strconv.Atoi(strings.TrimSuffix(day, "\n"))
	if newErr != nil {
		log.Fatal(newErr)
	}
	fmt.Println()
	return result
}

func inputMonth() int {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a month: ")
	month, err := scanner.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	result, newErr := strconv.Atoi(strings.TrimSuffix(month, "\n"))
	if newErr != nil {
		log.Fatal(newErr)
	}
	return result
}

func inputYear() int {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a year: ")
	year, err := scanner.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	result, newErr := strconv.Atoi(strings.TrimSuffix(year, "\n"))
	if newErr != nil {
		log.Fatal(newErr)
	}
	return result
}

func displayDate(d Date) string {
	return strconv.Itoa(d.year) + "/" + strconv.Itoa(d.month) + "/" + strconv.Itoa(d.day)
}

func calcDays(d Date) int {
	difference := d.year - 2000
	result := ((difference * 360) + ((d.month - 1) * 30) + (d.day - 1))
	return result
}
