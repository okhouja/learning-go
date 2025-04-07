package note

import (
	"fmt"
	"os"
	"time"
)

type Note struct {
	Title   string
	Content string
}

func New(title, content string) Note {
	return Note{
		Title:   title,
		Content: content,
	}
}

func (n Note) Display() {
	fmt.Println("Note Title:", n.Title)
	fmt.Println("Note Content:", n.Content)
}

func (n Note) Save() error {
	filename := fmt.Sprintf("note_%d.txt", time.Now().Unix())
	return os.WriteFile(filename, []byte(n.Title+"\n"+n.Content), 0644)
}