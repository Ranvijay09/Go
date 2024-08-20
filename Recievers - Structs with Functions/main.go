package main

import (
	"fmt"
	"time"
)

type Employee struct {
	FirstName string
	LastName  string
	Gender    string
	DOJ       time.Time
}

func (e *Employee) printDetails() {
	fmt.Println(*e)
}

func main() {
	var emp1 Employee

	emp1.FirstName = "Missy"
	emp1.LastName = "Cooper"
	emp1.Gender = "Female"

	emp2 := Employee{
		FirstName: "Sheldon",
		LastName:  "Cooper",
		Gender:    "Male",
	}

	emp1.printDetails()

	emp2.printDetails()
}
