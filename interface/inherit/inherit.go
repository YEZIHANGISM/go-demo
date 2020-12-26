package inherit

import (
	"log"
)

type Task struct {
	Command string
	*log.Logger
}

func TaskFactory(command string, logger *log.Logger) *Task {
	return &Task{command, logger}
}

func Inherit() {
}
