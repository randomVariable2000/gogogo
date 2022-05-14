package main

import "fmt"

func main() {
	var n int

	defer func() {
		// what will Println?
		fmt.Println("Defer 1:", n)
	}()

	n = 3

	func() {
		fmt.Println("Normal :", n)
	}()
}
