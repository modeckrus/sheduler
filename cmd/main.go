package main

import (
	"fmt"

	"github.com/modeckrus/sheduler/task"
)

type Message struct {
	Str string
}

func Task(message Message) {

}

func main() {
	meta, err := task.NewFuncRegistry().Add(Task)
	fmt.Println(meta, err)
}
