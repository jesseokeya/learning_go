package main

import "testing"

func TestRemoveEntryByID(t *testing.T) {
	length := len(a.GetAllEntries())
	newEntry := Entry{
		length + 1,
		"TestFirstName",
		"TestLastName",
		"TestAddress",
		"TestPhoneNumber",
	}
	a.AddEntry(newEntry)
	a.RemoveEntryByID(length + 1)
	a.AddEntry(newEntry)
	getAllEntries := a.GetAllEntries()
	if getAllEntries[len(getAllEntries)-1].ID != newEntry.ID {
		t.Error("something went wrong in removing an entry by id")
	}
}

func TestRemoveEntry(t *testing.T) {
	length := len(a.GetAllEntries())
	newEntry := Entry{
		length + 1,
		"TestFirstName",
		"TestLastName",
		"TestAddress",
		"TestPhoneNumber",
	}
	a.AddEntry(newEntry)
	getAllEntries := a.GetAllEntries()
	entryRemoved := a.RemoveEntry("TestFirstName", "TestLastName", " TestAddress", "TestPhoneNumber")
	if getAllEntries[len(getAllEntries)-1].FirstName != entryRemoved[0].FirstName {
		t.Error("something went wrong in removing an Entry")
	}
}

func TestSearch(t *testing.T) {
	length := len(a.GetAllEntries())
	newEntry := Entry{
		length + 1,
		"Jesse",
		"James",
		"944 Dollar Sign Avenue Ottawa, ON",
		"61398999450",
	}
	a.AddEntry(newEntry)
	result := a.Search("Jesse", "James", "61398999450")
	if result.FirstName != newEntry.FirstName {
		t.Error("something went wrong in searching for an Entry")
	}
	secondSearch := a.Search("__*^^$#@$$!@$$($)")
	validate := Entry{}
	if secondSearch != validate {
		t.Error("something went wrong in searching for an Entry")
	}
}

func TestEditEntry(t *testing.T) {
	length := len(a.GetAllEntries())
	a.EditEntry(length-1, Entry{})
	a.RemoveEntry("Jesse", "TestFirstName")
}
