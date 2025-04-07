package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/todo_note_app/note"
	"example.com/todo_note_app/todo"
)

type Saveable interface {
	Save() error
	Display()
}

// 💡 Generic function that takes any type that implements Saveable
func processAndSave[T Saveable](item T) {
	item.Display()
	err := item.Save()
	if err != nil {
		fmt.Println("❌ Error saving:", err)
		return
	}
	fmt.Println("✅ Saved successfully!")
}

func getUserInput(prompt string) string {
	fmt.Print(prompt + ": ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	noteTitle := getUserInput("Enter note title")
	noteContent := getUserInput("Enter note content")
	n := note.New(noteTitle, noteContent)
	processAndSave(n)

	todoText := getUserInput("Enter todo task")
	t := todo.New(todoText)
	processAndSave(t)
}
