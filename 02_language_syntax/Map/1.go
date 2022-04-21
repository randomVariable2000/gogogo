// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

import "fmt"

// Add imports.

func main() {

	// Declare and make a map of integer type values.
	type mapType map[int8]int
	map1 := make(mapType)

	// Initialize some data into the map.
	map1[1] = 10086
	map1[2] = 1008611
	map1[3] = 12356

	// Display each key/value pair.
	for k, v := range map1 {
		fmt.Println(k, v)
	}
}
