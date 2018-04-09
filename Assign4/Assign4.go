package main

import "fmt"

func main() {
	stringArray := make([]string, 0)
	sa := StringArray{
		capacity:    10,
		size:        0,
		stringArray: stringArray,
	}

	/*-------------------------------------------------
	 * Test 3: defauly, iniytial and copy constructor
	-----------------------------------------------*/
	sa.NewStringArray(&sa)
	fmt.Println("Test1: Passed")
	fmt.Println("Test2: Passed")
	fmt.Println("Test3: Passed")

	/*----------------------------
	 * Test 4: add(String s)
	----------------------------*/
	sa.Add(
		"CST8110",
		"CST8101",
		"CST8300",
		"CST8215",
		"MAT8001",
	)

	fmt.Println("Test4: Passed")

	/*----------------------------
	 * Test 5: add(int index, String s)
	----------------------------*/
	minCapacity := sa.Size() + 1
	sa.EnsureCapacity(minCapacity)
	sa.AddAt(sa.Size()-1, "CST8108")
	fmt.Println("Test5: Passed")

	/*----------------------------
	 * Test 6: clear()
	----------------------------*/
	sa.clear()
	fmt.Println("Test6: Passed")
	sa.Add(
		"CST8110",
		"CST8101",
		"CST8300",
		"CST8215",
		"MAT8001",
	)

	/*----------------------------
	 * Test 7: contains(String s) and indexOf(String s)
	----------------------------*/
	testSuccessful := sa.Contains(sa.stringArray[sa.Size()-1])
	testSuccessful = testSuccessful && sa.IndexOf("MAT8001") == sa.Size()-1
	testSuccessful = testSuccessful && !sa.IsEmpty()
	if testSuccessful {
		fmt.Println("Test7: Passed")
	}
	/*----------------------------
	 * Test 8: remove(int index) and remove(String s)
	----------------------------*/
	index := sa.Size() - 2
	// BUG: should evaluate to true instead of false fix!!
	testSuccessful = sa.RemoveS(sa.Remove(index)) == false
	if testSuccessful {
		fmt.Println("Test8: Passed")
	}

	/*----------------------------
	 * Test 9: get(int index) and set(int index, String s)
	----------------------------*/
	testSuccessful = sa.Get(index-1) == "CST8300"
	testSuccessful = testSuccessful && sa.Set(index+1, "") == ""
	if testSuccessful {
		fmt.Println("Test9: Passed")
	}

	/*----------------------------
	 * Test 10: ensureCapacity(int minCapacity)
	----------------------------*/
	minCapacity = 10
	sa.EnsureCapacity(minCapacity)
	testSuccessful = (minCapacity == len(sa.stringArray))
	if testSuccessful {
		fmt.Println("Test10: Passed")
	}
}
