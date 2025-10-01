package tui

import (
	"fmt"
	"tasker/internal/gen"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

type generationMsg struct {
	botResponse string
	err         error
}

func performGeneration(userMessage string) tea.Cmd {
	return func() tea.Msg {
		botResponse, err := gen.GenerateTask(userMessage)
		if err != nil {
			fmt.Println(err)
			return generationMsg{err: err}
		}
		botResponse, err = glamour.Render(botResponse, "dark")

		return generationMsg{botResponse: botResponse, err: err}
	}
}
