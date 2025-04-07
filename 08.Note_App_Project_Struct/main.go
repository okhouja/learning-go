package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main(){
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}
	fmt.Println("Note saved successfully!")
}

func getNoteData() (string, string){
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string{
	fmt.Printf("%v ",prompt)
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