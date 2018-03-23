package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	a       AddressBook
	message string
)

func init() {
	a = AddressBook{}
	fmt.Println(" \n------------------------------------ ")
	fmt.Println(" 📕 Welcome To Your Address Book 📗 ")
	fmt.Println("------------------------------------ ")

	fmt.Println("😌 Here Are Your Options Today 😌 ")
	fmt.Println("------------------------------------ ")
	fmt.Println(" 1 ⇢ Load from file ")
	fmt.Println(" 2 ⇢ Save to file")
	fmt.Println(" 3 ⇢ Add an entry")
	fmt.Println(" 4 ⇢ Remove an entry")
	fmt.Println(" 5 ⇢ Edit an existing entry")
	fmt.Println(" 6 ⇢ Search for a specific entry")
	fmt.Println(" 7 ⇢ Sort the address book")
	fmt.Println(" 8 ⇢ Display data in address book")
	fmt.Println(" 9 ⇢ Quit")
	fmt.Println(" ----------------------------------- ")
}

func main() {
	fmt.Println(" Choose From The Numbers Above ")
	fmt.Println(" ----------------------------------- ")
	for {
		scanner := bufio.NewReader(os.Stdin)
		fmt.Print(" Enter Any Valid Operation 1 - 9: ")
		operation, _ := scanner.ReadString('\n')
		operation = strings.TrimSuffix(operation, "\n")
		convOperation, err := strconv.Atoi(operation)
		if err != nil {
			log.Fatal(err)
		}
		switch convOperation {
		case 1:
			break
		case 2:
		case 3:
			break
		case 4:
			break
		case 5:
			break
		case 6:
			break
		case 7:
			break
		case 8:
			break
		case 9:
			return
		default:
			break
		}
	}

}

func prettyPrint(a []Entry) {
	fmt.Println(" \n----------------------------------- ")
	fmt.Println(" All Data Available In Address Book ")
	fmt.Println("------------------------------------ ")
	for _, entry := range a {
		fmt.Printf("Id: %v\n", entry.ID)
		fmt.Println("FirstName: " + entry.FirstName)
		fmt.Println("LastName: " + entry.LastName)
		fmt.Println("PhoneNumber: " + entry.PhoneNumber)
		fmt.Println("------------------------------------ ")
	}

}

func displayAllData() {
	data := a.GetAllEntries()
	prettyPrint(data)
}
