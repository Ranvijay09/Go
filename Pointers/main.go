package main

import "fmt"

func main() {
	var myString string

	myString = "Red"

	fmt.Println("My String set to", myString)

	changeValueUsingPointer(&myString)

	fmt.Println("My String changed to", myString)
}

func changeValueUsingPointer(str *string) {
	fmt.Println("Pointer Address for myString:", str)
	*str = "Greeen"
}
