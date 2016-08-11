package main

import "fmt"

func main() {
	// intialize a counter
	number := 0
	for i := 0; i <= 1000; i++ {
		if i%3 == 0 { // divisble by 3?
			number += i // add to counter
		} else if i%5 == 0 { // divisible by 5?
			number += i // add it to the counter
		}
	}
	fmt.Println(number) // what number did the counter get to?
}
