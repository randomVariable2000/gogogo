// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

import "fmt"

// Add imports.

func main() {

	// Declare a nil slice of integers.
	var slice1 []int

	// Append numbers to the slice.
	slice1 = append(slice1, 10086)

	// Display each value in the slice.
	for _, i2 := range slice1 {
		fmt.Println(i2)
	}

	// Declare a slice of strings and populate the slice with names.
	slice2 := []string{"huang", "chen", "li", "zhao"}

	// Display each index position and slice value.
	for i, s := range slice2 {
		fmt.Println(i, s)
	}

	// Take a slice of index 1 and 2 of the slice of strings.
	slice3 := slice2[1:3]

	// Display each index position and slice values for the new slice.
	for i, s := range slice3 {
		fmt.Println(i, s)
	}
}
