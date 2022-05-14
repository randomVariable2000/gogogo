package main

import (
	"crypto/sha1"
	"fmt"
	"strconv"
	"sync"
)

// Sample program to show how the routine scheduler
// will time slice Goroutines on single thread.
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Create Goroutines")

	go func() {
		printHashes("A")
		wg.Done()
	}()

	go func() {
		printHashes("B")
		wg.Done()
	}()

	wg.Wait()
}

func printHashes(prefix string) {
	for i := 1; i <= 50000; i++ {
		num := strconv.Itoa(i)
		sum := sha1.Sum([]byte(num))
		fmt.Printf("%s: %05d: %x\n", prefix, i, sum)
	}

	fmt.Println("Completed", prefix)
}
