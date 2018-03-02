package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

func testDouble(t *testing.T) {
	values := []string{"23.5", "45.7", "67.9", "98.1", "23.5", "12.00", "13.000332"}
	for _, num := range values {
		val := double(num)
		if reflect.TypeOf(val).String() != "float64" {
			t.Error("Invalid Data Type, Must Always Be float64")
		}
	}
}

func testHandleErr(t *testing.T) {
	price, err := decimal.NewFromString("136.02")
	fmt.Println(price)
	if err != nil {
		HandleErr(err)
	} else {
		fmt.Println(price)
	}
}
