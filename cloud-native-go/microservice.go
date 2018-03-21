package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	httplogger "github.com/jesseokeya/go-httplogger"
	"github.com/jesseokeya/learning_go/cloud-native-go/api"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", index)
	r.HandleFunc("/api/echo", api.EchoHandleFunc)
	r.HandleFunc("/api/hello", api.HelloHandleFunc)

	r.HandleFunc("/api/books", api.BooksHandleFunc)
	r.HandleFunc("/api/books/", api.BookHandleFunc)

	http.ListenAndServe(port(), httplogger.Golog(r))
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to Cloud Native Go.")
}
