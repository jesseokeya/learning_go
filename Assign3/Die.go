package main

import (
	"fmt"
	"math/rand"
)

// Die struct holds all methods applied to the die object
type Die struct {
	dieValue int
}

func (d *Die) rollDie() {
	d.dieValue = rand.Intn(6) + 1
}

func (d *Die) display() {
	fmt.Print(d.dieValue)
}

func (d *Die) getValue() int {
	return d.dieValue
}
