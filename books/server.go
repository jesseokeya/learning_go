package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/getAllBooks", getAllBooks)
	http.ListenAndServe(port("3000"), nil)
}

func allBooks() []Book {
	books := []Book{
		{Title: "Narnia", Author: "Jame Paul", ISBN: "01234537"},
		{Title: "Harry Porter", Author: "Jesse Okeya", ISBN: "01234569"},
		{Title: "Batman", Author: "Sith Hernderson", ISBN: "012345678"},
	}
	return books
}

func port(num string) string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = num
	}
	fmt.Println("server running on port *" + port)
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome To The Books Api")
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/jsonl; charset=utf-8")
	booksInJSON, err := json.Marshal(allBooks())
	if err != nil {
		panic(err)
	}
	w.Write(booksInJSON)
}
