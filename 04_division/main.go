package main

import "fmt"

func main() {
	// The big number
	var b int
	fmt.Println("Please enter a big number, then hit Enter.")
	fmt.Scan(&b)
	// The little number
	var l int
	fmt.Println("Please enter a little number, then hit Enter.")
	fmt.Scan(&l)
	fmt.Println("This is the remainder of ", b, "%", l, "=", b%l)
}
