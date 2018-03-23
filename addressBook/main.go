package main

import "fmt"

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
	fmt.Println(" \n----------------------------------- ")
	fmt.Println(" All Data Available In Address Book ")
	fmt.Println("------------------------------------ ")
	//fmt.Println(a.GetAllEntries())
}

func main() {
	data := a.GetAllEntries()
	prettyPrint(data)
}

func prettyPrint(a []Entry) {
	for _, entry := range a {
		fmt.Printf("Id: %v\n", entry.ID)
		fmt.Println("FirstName: " + entry.FirstName)
		fmt.Println("LastName: " + entry.LastName)
		fmt.Println("PhoneNumber: " + entry.PhoneNumber)
		fmt.Println("------------------------------------ ")
	}
}
