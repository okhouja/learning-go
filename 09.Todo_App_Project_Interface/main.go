package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

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
	err = saveData(todo)

	if err != nil {
		fmt.Println("Saving the todo failed.")
		return
	}

	userNote.Display()
	err = saveData(userNote)

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving failed.")
		return err
	}

	fmt.Println("Saved successfully!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	// var value string
	// fmt.Scanln(&value) // good for one word input and not for multiple words

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n') // good for multiple words input

	if err != nil {
		fmt.Println(err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
