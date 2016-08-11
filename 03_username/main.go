package main

import "fmt"

func main() {
	fmt.Println("Please enter your name, then hit Enter.")
	var name string
	fmt.Scanf("%s", &name)
	fmt.Println("Hello", name, "!")
}
