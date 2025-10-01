package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
		spCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			if !m.isLoading {
				userMessage := m.textarea.Value()
				if userMessage == "" {
					return m, nil
				}

				m.isLoading = true
				m.response = ""
				m.textarea.Reset()
				m.viewport.GotoBottom()

				return m, tea.Batch(performGeneration(userMessage), m.spinner.Tick)
			}
		}

	case generationMsg:
		m.isLoading = false
		m.response = m.textStyle.Render(msg.botResponse)

		if msg.err != nil {
			m.response = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render(fmt.Sprintf("Error: %s", msg.err.Error()))
		}
		m.viewport.SetContent(m.response)
		m.viewport.GotoBottom()

		return m, nil

	case spinner.TickMsg:
		if m.isLoading {
			m.spinner, spCmd = m.spinner.Update(msg)
			return m, tea.Batch(tiCmd, vpCmd, spCmd)
		}

	case tea.WindowSizeMsg:
		m.viewport.Height = msg.Height - 3
		m.viewport.Width = msg.Width
		m.textarea.SetWidth(msg.Width)

		m.viewport.SetContent(m.response)
	}

	return m, tea.Batch(tiCmd, vpCmd)
}
