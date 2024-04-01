package main

import (
	"fmt"
	"go_tutorials/structs/user"
)

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthday (DD/MM/YYYY): ")

	appUser, err := user.NewUser(userFirstName, userLastName, userBirthDate)
	handleError(err)

	admin, err := user.NewAdmin("test@exapmle.com", "test123")
	handleError(err)

	admin.OutputUserData()
	appUser.OutputUserData()
	appUser.ClearUserName()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
