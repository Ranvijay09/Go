package main

import "fmt"

var myString string //package level variable

func main() {
	fmt.Println("Hello, Vinod")

	var greetMsg string

	greetMsg = "Good Morning!"

	fmt.Println(greetMsg)

	var i int

	i = 2

	fmt.Println("Integer i set to", i)

	str1, str2 := saySomething()

	fmt.Println("Function returned", str1, str2)
}

func saySomething() (string, string) {
	return "Something", "Something!"
}
