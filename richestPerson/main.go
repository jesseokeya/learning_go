package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("./src"))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(port("8000"), nil))
}

func port(p string) string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = p
	}
	fmt.Println("server running on port *" + port)
	return ":" + port
}
