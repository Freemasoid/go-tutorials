package main

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func newUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birthday are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthday (DD/MM/YYYY): ")

	appUser, err := newUser(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.outputUserData()
	appUser.clearUserName()
}

func (u User) outputUserData() {
	// ...
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
