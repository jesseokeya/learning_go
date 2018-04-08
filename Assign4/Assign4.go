package main

import "fmt"

func main() {
	stringArray := make([]string, 0)
	sa := StringArray{
		capacity:    10,
		size:        0,
		stringArray: stringArray,
	}

	sa.NewStringArray(&StringArray{})
	sa.Add("Jesse", "Naomi", "Peter", "Wee", "hanna")
	sa.AddAt(2, "James")
	fmt.Println(sa.Contains("James"))
	fmt.Println(sa.Remove(1))
	fmt.Println(sa.stringArray)
}
