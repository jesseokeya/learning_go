package main

import (
	"reflect"
	"strconv"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOurDate(t *testing.T) {
	d1 := Date{
		31,
		12,
		2017,
	}
	d2 := OurDate(31, 12, 2017)
	Describe("Comparing Two Dates Structs", func() {
		Context("Two dates should be equal", func() {
			It("should be equal", func() {
				Expect(d1.day).To(Equal(d2.day))
				Expect(d1.month).To(Equal(d2.month))
				Expect(d1.year).To(Equal(d2.year))
			})
		})
	})
}

func TestDisplayDate(t *testing.T) {
	input := Date{
		1,
		1,
		2018,
	}
	result := displayDate(input)
	Describe("Check If Display Date Returns Proper String", func() {
		Context("Compare Strings", func() {
			It("should be equal", func() {
				answer := strconv.Itoa(input.year) + "/" + strconv.Itoa(input.month) + "/" + strconv.Itoa(input.day)
				Expect(result).To(Equal(answer))
			})
		})
	})
}

func TestCalcDays(t *testing.T) {
	input := Date{
		31,
		12,
		2017,
	}
	d1 := calcDays(input)
	d2 := calcDays(input)
	Describe("Checking Accuracy Of calcDays Calculation", func() {
		Context("Result Should Be An Integer", func() {
			It("should be equal", func() {
				Expect(d1).To(Equal(d2))
				if reflect.TypeOf(d1).String() != "int" {
					t.Error("Invalid Data Type, Must Always Be An int")
				}
			})
		})
	})
}
