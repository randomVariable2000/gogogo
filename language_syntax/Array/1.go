// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

import "fmt"

// Add imports.

func main() {

	// Declare an array of 5 strings set to its zero value.
	var arr1 [5]string

	// Declare an array of 5 strings and pre-populate it with names.
	arr2 := [5]string{"wu", "chen", "li", "chai", "huang"}
	// Assign the populated array to the array of zero values.
	arr1 = arr2
	// Iterate over the first array declared.
	// Display the string value and address of each element.
	for i, s := range arr1 {
		fmt.Println(s, &arr1[i])
	}
}
