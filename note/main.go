package main

import (
	"bufio"
	"fmt"
	"go_tutorials/note/note"
	"go_tutorials/note/todo"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = todo.Save()

	if err != nil {
		fmt.Println("saving todo failed")
		return
	}

	fmt.Println("The todo is saved!")

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("saving the note failed")
		return
	}

	fmt.Println("The note is saved!")
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return title, content
}
