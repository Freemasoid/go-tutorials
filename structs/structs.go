package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthday (DD/MM/YYYY): ")

	var appUser = User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}

	outputUserData(&appUser)
}

func outputUserData(u *User) {
	// ...
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
