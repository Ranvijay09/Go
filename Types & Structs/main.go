package main

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	DOB       time.Time
}

func main() {
	user := User{
		FirstName: "Vinod",
		LastName:  "Pujari",
	}

	fmt.Println(user.FirstName, user.LastName)
	fmt.Println(user)
}
