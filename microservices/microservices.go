package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)
	http.ListenAndServe(port("3000"), nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome To Cloud Native Go")
}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}

func port(num string) string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = num
	}
	fmt.Println("server running on port *" + port)
	return ":" + port
}
