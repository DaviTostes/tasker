package tui

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func setViewportError(viewport viewport.Model, err error) {
	viewport.SetContent(
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("9")).
			Render(fmt.Sprintf("Error: %s", err.Error())),
	)
}

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
				m.textarea.Reset()
				m.viewport.GotoBottom()

				return m, tea.Batch(performGeneration(userMessage), m.spinner.Tick)
			}

		case tea.KeyCtrlY:
			clipboard.WriteAll(m.task.GetText())
		}

	case generationMsg:
		m.isLoading = false
		m.task = msg.task

		if msg.err != nil {
			setViewportError(m.viewport, msg.err)
			return m, nil
		}

		md, err := m.task.RenderMd()
		if err != nil {
			setViewportError(m.viewport, err)
			return m, nil
		}

		m.viewport.SetContent(md)

	case spinner.TickMsg:
		if m.isLoading {
			m.spinner, spCmd = m.spinner.Update(msg)
			return m, tea.Batch(tiCmd, vpCmd, spCmd)
		}

	case tea.WindowSizeMsg:
		m.viewport.Height = msg.Height - 4
		m.viewport.Width = msg.Width
		m.textarea.SetWidth(msg.Width)

		md, err := m.task.RenderMd()
		if err != nil {
			setViewportError(m.viewport, err)
			return m, nil
		}

		m.viewport.SetContent(md)
	}

	return m, tea.Batch(tiCmd, vpCmd)
}
