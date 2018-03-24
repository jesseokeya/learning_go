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
	fmt.Println(" 1 ⇢ Display data in address book")
	fmt.Println(" 2 ⇢ Save to file")
	fmt.Println(" 3 ⇢ Add an entry")
	fmt.Println(" 4 ⇢ Remove an entry")
	// fmt.Println(" 5 ⇢ Edit an existing entry")
	// fmt.Println(" 6 ⇢ Search for a specific entry")
	// fmt.Println(" 7 ⇢ Sort the address book")
	fmt.Println(" 5 ⇢ Quit")
	fmt.Println(" ----------------------------------- ")
}

func main() {
	fmt.Println(" Choose From The Numbers Above ")
	fmt.Println(" ----------------------------------- ")
	for {
		scanner := bufio.NewReader(os.Stdin)
		fmt.Print(" Please choose what you'd like to do with the database (1 - 9): ")
		operation, _ := scanner.ReadString('\n')
		operation = strings.TrimSuffix(operation, "\n")
		convOperation, err := strconv.Atoi(operation)
		if err != nil {
			log.Fatal(err)
		}
		switch convOperation {
		case 1:
			displayAllData()
			fmt.Println()
			break
		case 2:
			newEntry := createNewEntry()
			a.AddEntry(newEntry)
		case 3:
			newEntry := createNewEntry()
			a.AddEntry(newEntry)
		case 4:
			remove := getUserInput(" Remove Existing query using FirstName, LastName, Address or PhoneNumber: ")
			dataRemoved := a.RemoveEntry(remove)
			fmt.Println(" ⛑ The Data Below was successfully deleted ⛑ ")
			prettyPrint(dataRemoved)
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
			fmt.Println(" Quitting Address Book.. ")
			return
		default:
			break
		}
	}

}

func prettyPrint(a []Entry) {
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

func createNewEntry() Entry {
	index := len(a.GetAllEntries())
	fmt.Println(" 🗒 Create a new addressBook entry 🗒 ")
	fname := getUserInput(" Enter First Name: ")
	lname := getUserInput(" Enter Last Name: ")
	address := getUserInput(" Enter Address: ")
	phone := getUserInput(" Enter Phone Number: ")
	fmt.Println()
	return Entry{index + 1, fname, lname, address, phone}
}

func getUserInput(message string) string {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Print(message)
	input, _ := scanner.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	return input
}
