package main

import "testing"

var (
	sa = StringArray{0, 10, []string{}}
)

func init() {
	sa.NewStringArray(&StringArray{})
}

func TestAdd(t *testing.T) {
	allCourses := []string{
		"CST8101",
		"CST8110",
		"CST8215",
		"CST8300",
		"MAT8001",
		"empty",
	}

	for _, course := range allCourses {
		sa.Add(course)
	}
	if sa.IsEmpty() {
		t.Error("String Array Object Should Not Be Empty ")
	}
}

func TestAddAt(t *testing.T) {
	capacity := sa.Capacity()
	sa = StringArray{10, capacity, sa.stringArray}
	sa.AddAt(29, "CST8108")
	if sa.IsEmpty() {
		t.Error("String Array Object Should Not Be Empty ")
	}
}

func TestClear(i *testing.T) {
	sa.clear()
}
