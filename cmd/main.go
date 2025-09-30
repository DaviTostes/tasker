package main

import (
	"log"
	"tasker/internal/model"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(model.InitialModel())
	if _, err := p.Run(); err != nil {
		log.Fatalln(err)
	}
}
