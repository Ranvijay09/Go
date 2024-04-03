package main

import "fmt"

func main() {
	//Array Definations
	arr1 := [2]int{2} //Fixed Size
	var arr2 = [3]int{2, 4, 8}

	arr1[1] = 8 //Can cahnge data using index number
	arr2[1] = 7 //After printing array, the sequence will be sam as per index

	fmt.Println(arr1, arr2)

	//Slice Definations
	slice1 := []int{}

	slice1 = append(slice1, 6)
	slice1 = append(slice1, 2, 3, 5) //Slices are growable, you can append 1 or more elements at a time using append()

	slice1[1] = 4 // can change data using index number only if index lies between existing slice length

	fmt.Println(slice1)

}
