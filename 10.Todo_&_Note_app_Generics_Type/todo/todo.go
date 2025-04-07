package todo

import (
	"fmt"
	"os"
	"time"
)

type Todo struct {
	Task string
	Done bool
}

func New(task string) Todo {
	return Todo{Task: task, Done: false}
}

func (t Todo) Display() {
	fmt.Println("Todo Task:", t.Task)
}

func (t Todo) Save() error {
	filename := fmt.Sprintf("todo_%d.txt", time.Now().Unix())
	return os.WriteFile(filename, []byte(t.Task), 0644)
}
