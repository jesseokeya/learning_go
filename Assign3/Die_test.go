package main

import (
	"reflect"
	"testing"
)

var die = Die{0}

func TestRollDie(t *testing.T) {
	die.rollDie()
	firstRoll := die.getValue()
	die.rollDie()
	secondRoll := die.getValue()
	if firstRoll == secondRoll {
		t.Error("roll die function should not give you the same values ")
	}
}

func TestDisplay(t *testing.T) {
	die.display()
}

func TestGetValue(t *testing.T) {
	die.rollDie()
	value := die.getValue()
	if reflect.TypeOf(value).String() != "int" {
		t.Error("getValue() result should be type of int ")
	}
}
