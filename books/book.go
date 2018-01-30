package main

import (
	"encoding/json"
)

// Book type with Name, Author and ISBN
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
	// define the book
}

// ToJSON to be used for mashalling of Book type
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

// FromJSON to bes used for unmarshalling of Book type
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

// func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {

// }
