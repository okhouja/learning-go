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

// type displayer interface {
// 	Display()
// }

// type outputtable interface {
// 	Save()
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

func main() {
// try any type interfaces
	printSomething("Hello")
	printSomething(1)
	printSomething(1.1)


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

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)

}

// Any type interfaces
func printSomething(value interface{}) {
	// switch type
	switch value.(type) {
	case int:
		fmt.Println("Integer",value)
	case float64:
		fmt.Println("Float",value)
	case string:
		fmt.Println("String")
	default:
		// .....
		fmt.Println("Unknown")
	}
	fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
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
