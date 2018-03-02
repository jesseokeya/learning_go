package main

import (
	"reflect"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDateCalculator(t *testing.T) {
	DateCalculator()
}

func TestInputDate(t *testing.T) {
	input := Date{
		31,
		12,
		2017,
	}
	result := inputDate(input)
	Describe("Comparing Two Dates Structs", func() {
		Context("Two dates should be equal", func() {
			It("should be equal", func() {
				Expect(input.day).To(Equal(result.day))
				Expect(input.month).To(Equal(result.month))
				Expect(input.year).To(Equal(result.year))
			})
		})
	})
}

func TestDisplay(t *testing.T) {
	input := display("This Is A Test")
	Describe("Checking If Display Returns Proper String", func() {
		Context("Two Strings Check and DataType", func() {
			It("should be equal", func() {
				Expect(input).To(Equal("This Is A Test"))
			})
		})
	})
}

func TestCalculateDifference(t *testing.T) {
	input := []Date{
		{31, 12, 2017},
		{1, 1, 2018},
	}

	message, result := calculateDifference(input[0], input[1])
	Describe("Checking If Display Returns Proper Message And Calculation", func() {
		Context("Verifying Message And Difference Of Dates", func() {
			It("should be equal", func() {
				messageResult := "2017/12/31 is same as 2018/1/1"
				Expect(result).To(Equal(0))
				Expect(message).To(Equal(messageResult))
			})
		})
	})
}

func TestGenerateDisplay(t *testing.T) {
	differences := []int{0, 34, 20, 10, 5, 4, 2, -1, 10000}
	d := []Date{
		{31, 12, 2002},
		{1, 2, 2016},
	}
	for _, num := range differences {
		result := generateDisplay(d[0], d[1], num)
		if reflect.TypeOf(result).String() != "string" {
			t.Error("Invalid Data Type, Must Always Be A string")
		}
	}
}
