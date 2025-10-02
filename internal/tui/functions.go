package tui

import (
	"context"
	"fmt"
	"tasker/internal/gen"

	tea "github.com/charmbracelet/bubbletea"
)

type generationMsg struct {
	task gen.Task
	err  error
}

func performGeneration(userMessage string) tea.Cmd {
	return func() tea.Msg {
		task, err := gen.TaskFlow.Run(context.Background(), userMessage)
		if err != nil {
			fmt.Println(err)
			return generationMsg{err: err}
		}

		return generationMsg{task: task, err: err}
	}
}
