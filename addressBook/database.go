package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// DB is an interface that interacts with the addressBook database
type DB struct{}

func (d *DB) create(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
}

func (d *DB) writeToFile(location string, data string) {
	err := ioutil.WriteFile(location, []byte(data), 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func (d *DB) readFromFile(location string) []byte {
	byte, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatal(err)
	}
	return byte
}

// ByteToString converts an array of bytes to string
func ByteToString(b *[]byte) string {
	result := *b
	return string(result[:])
}

// DeleteFile removes file in storage
func (d *DB) DeleteFile(path string) {
	var err = os.Remove(path)
	if err != nil {
		panic(err)
	}
	fmt.Println("==> done deleting file -> " + path)
}
