package main

import (
	"testing"
)

var (
	database DB
	filename string
)

func init() {
	database = DB{}
	filename = "storage/testing.json"
}

func TestCreate(t *testing.T) {

	err := database.create(filename)
	if err != nil {
		t.Error("something went wrong in creating a file")
	}
}

func TestWriteToFile(t *testing.T) {
	length := len(a.GetAllEntries())
	newEntry := Entry{
		length + 1,
		"TestFirstName",
		"TestLastName",
		"TestAddress",
		"TestPhoneNumber",
	}
	a.AddEntry(newEntry)
	if length == len(a.GetAllEntries()) {
		t.Error("something went wrong in writing to a file")
	}
}

func TestDeleteFile(t *testing.T) {
	path := filename
	err := database.DeleteFile(path)
	if err != nil {
		t.Error("something went wrong in deleting a file")
	}
}
