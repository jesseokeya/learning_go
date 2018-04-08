package main

// StringArray initializes the object
type StringArray struct {
	capacity    int
	size        int
	stringArray []string
}

// NewStringArray creates a new StringArray class
func (sa *StringArray) NewStringArray(copy *StringArray) {
	if copy.capacity == 0 && copy.size == 0 && len(copy.stringArray) == 0 {
		sa.capacity = 10
		sa.size = 0
	} else {
		sa.capacity = copy.capacity
		sa.size = copy.size
		sa.stringArray = copy.stringArray
	}
}

// Add Appends the specified string to the end of this StringArray. Increases capacity if needed.
func (sa *StringArray) Add(s ...string) {
	if sa.size != 0 {
		sa.EnsureCapacity(sa.size + 1)
		sa.stringArray = make([]string, sa.size)
		sa.stringArray = s
	} else {
		sa.size = sa.size + len(s)
		sa.EnsureCapacity(sa.size)
		count := 0
		for i := 0; i < len(sa.stringArray); i++ {
			if sa.stringArray[i] == "" {
				sa.stringArray[i] = s[count]
				count++
			}
		}
	}
	sa.size = len(sa.stringArray)
}

// AddAt inserts the specified string at the specified position in this StringArray.
// Increases capacity if needed.
func (sa *StringArray) AddAt(index int, s string) {
	if index < sa.size {
		sa.stringArray[index] = s
	} else {
		sa.size++
		sa.EnsureCapacity(sa.size)
		sa.stringArray[sa.size-1] = s
	}
}

// Capacity returns the capacity of this StringArray.
func (sa *StringArray) Capacity() int {
	return sa.capacity
}

// Returns the capacity of this StringArray.
func (sa *StringArray) clear() {
	sa.stringArray = make([]string, 0)
}

// Contains returns true if this StringArray contains the specified String.
func (sa *StringArray) Contains(s string) bool {
	result := false
	for _, stringValue := range sa.stringArray {
		if stringValue == s {
			result = true
		}
	}
	return result
}

// EnsureCapacity to ensure that it can hold at least the number of
// Strings specified by the minimum capacity argument.
func (sa *StringArray) EnsureCapacity(minCapacity int) {
	oldStringArray := sa.stringArray
	if minCapacity > sa.capacity {
		sa.stringArray = make([]string, minCapacity)
		for index, stringValue := range oldStringArray {
			sa.stringArray[index] = stringValue
			sa.size = len(sa.stringArray)
		}
	} else {
		sa.stringArray = make([]string, sa.size)
		for index, stringValue := range oldStringArray {
			sa.stringArray[index] = stringValue
			sa.size = len(sa.stringArray)
		}
	}
}

// Get returns the String at the specified position in this StringArray.
func (sa *StringArray) Get(index int) string {
	result := ""
	if sa.stringArray[index] != "" {
		result = sa.stringArray[index]
	}
	return result
}

// IndexOf returns the index of the first occurrence of the specified String in this StringArray,
// or -1 if this StringArray does not contain the String.
func (sa *StringArray) IndexOf(s string) int {
	result := -1
	for index, stringValue := range sa.stringArray {
		if s == stringValue {
			result = index
		}
	}
	return result
}

// IsEmpty returns true if this StringArray contains no Strings.
func (sa *StringArray) IsEmpty() bool {
	result := true
	if len(sa.stringArray) >= 0 {
		count := 0
		for _, stringValue := range sa.stringArray {
			if stringValue == "" {
				count++
			}
		}
		return count == len(sa.stringArray)
	}
	return result
}

// Remove removes the String at the specified position in this StringArray.
// Return the String previously at the specified position
func (sa *StringArray) Remove(index int) string {
	removed := ""
	if index < sa.size {
		removed = sa.stringArray[index]
		sa.stringArray[index] = ""
	}
	return removed
}

// RemoveS removes the first occurrence of the specified String from this StringArray, if it is present.
// Return true if this StringArray contained the specified String.
func (sa *StringArray) RemoveS(s string) bool {
	removed := false
	for index, stringValue := range sa.stringArray {
		if s == stringValue {
			sa.stringArray[index] = ""
			removed = true
		}
	}
	return removed
}

// Set replaces the String at the specified position in this StringArray with the specified String.
// Return the String previously at the specified position. Increases capacity if needed.
func (sa *StringArray) Set(index int, s string) string {
	previous := ""
	if index > sa.size {
		sa.size++
		sa.EnsureCapacity(sa.size)
		sa.stringArray[sa.size-1] = s
	} else {
		sa.stringArray[index] = s
	}
	return previous
}

// Size returns the number of Strings in this StringArray.
func (sa *StringArray) Size() int {
	return sa.size
}
