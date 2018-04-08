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
	sa.Add("Josh", "Jesse", "Naomi", "Peter", "Kaine", "Hanna", "Bekimbo", "Thomas")
	sa.AddAt(2, "James")
	fmt.Println(sa.Contains("James"))
	fmt.Println(sa.Remove(1))
	fmt.Println(sa)
}
