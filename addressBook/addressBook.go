package main

import (
	"encoding/json"
	"log"
)

// AddressBook holds data and implements logic for people
type AddressBook struct{}

var (
	location string
	db       DB
)

// Entry struct is a blueprint of how the addressBook Schema should be
type Entry struct {
	ID          int    `json:"_id,omitempty"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	Address     string `json:"address,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

func init() {
	location = "storage/address_book.json"
	db = DB{}
}

// LoadFromFile returns all data previously saved in the database
func (a *AddressBook) LoadFromFile() string {
	data := db.readFromFile(location)
	return ByteToString(&data)
}

// SaveToFile saves addressBook data to the database
func (a *AddressBook) SaveToFile(entry []Entry) {
	data, err := json.Marshal(entry)
	if err != nil {
		log.Fatal(err)
	}
	db.writeToFile(location, ByteToString(&data))
}

// AddEntry adds a new address_book to the database
func (a *AddressBook) AddEntry(entry Entry) {
	newData := []Entry{}
	data := a.LoadFromFile()
	err := json.Unmarshal([]byte(data), &newData)
	if err != nil {
		log.Fatal(err)
	}
	newData = append(newData, entry)
	a.SaveToFile(newData)
}

// RemoveEntry removes entry from AddressBook
func (a *AddressBook) RemoveEntry(query ...string) {
	data := []Entry{}
	load := a.LoadFromFile()
	err := json.Unmarshal([]byte(load), &data)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range data {
		for _, q := range query {
			if entry.FirstName == q || entry.LastName == q || entry.Address == q || entry.PhoneNumber == q {
				deleteEntry(entry, &data)
			}
		}
	}
	a.SaveToFile(data)
}

// RemoveEntryByID removes entry from AddressBook by id
func (a *AddressBook) RemoveEntryByID(id int) {
	data := []Entry{}
	load := a.LoadFromFile()
	err := json.Unmarshal([]byte(load), &data)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range data {
		if entry.ID == id {
			deleteEntry(entry, &data)
		}
	}
	a.SaveToFile(data)
}

// EditEntry finds an updates an entry
func (a *AddressBook) EditEntry(id int, updated Entry) {
	allEntries := a.GetAllEntries()
	for _, entry := range allEntries {
		if entry.ID == id {
			entry = updated
		}
	}
}

// GetAllEntries gets all Entries from the database
func (a *AddressBook) GetAllEntries() []Entry {
	result := []Entry{}
	allEntries := a.LoadFromFile()
	err := json.Unmarshal([]byte(allEntries), &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// Search searches entry if found in AddressBook
func (a *AddressBook) Search(query ...string) Entry {
	data := []Entry{}
	load := a.LoadFromFile()
	err := json.Unmarshal([]byte(load), &data)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range data {
		for _, q := range query {
			if entry.FirstName == q || entry.LastName == q || entry.Address == q || entry.PhoneNumber == q {
				return entry
			}
		}
	}
	return Entry{}
}

func deleteEntry(data Entry, entry *[]Entry) {
	newEntry := []Entry{}
	for _, entry := range *entry {
		if data.ID != entry.ID {
			newEntry = append(newEntry, entry)
		}
	}
	*entry = newEntry
}
