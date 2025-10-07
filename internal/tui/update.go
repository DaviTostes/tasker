package tui

import (
	"fmt"

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

type batch struct {
	tiCmd     tea.Cmd
	vpCmd     tea.Cmd
	spCmd     tea.Cmd
	editorCmd tea.Cmd
	cmds      []tea.Cmd
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	b := batch{}

	if m.isEditing {
		m.editor, b.editorCmd = m.editor.Update(msg)
		b.cmds = append(b.cmds, b.editorCmd)
	} else {
		m.textarea, b.tiCmd = m.textarea.Update(msg)
		m.viewport, b.vpCmd = m.viewport.Update(msg)
		b.cmds = append(b.cmds, b.tiCmd, b.vpCmd)
	}

	if m.isLoading {
		m.spinner, b.spCmd = m.spinner.Update(msg)
		b.cmds = append(b.cmds, b.spCmd)
	}

	keyMap := GetKeyMap(m, b)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		def := keyMap[msg.Type]
		if def != nil {
			return def()
		}

	case generationMsg:
		m.isLoading = false
		m.task = msg.task

		if msg.err != nil {
			setViewportError(m.viewport, msg.err)
			return m, nil
		}

		m.appContent = m.task.GetText()

		renderedContent, err := m.task.RenderMd()
		if err != nil {
			setViewportError(m.viewport, err)
			return m, nil
		}

		m.viewport.SetContent(renderedContent)
		m.viewport.GotoBottom()

	case spinner.TickMsg:
		if m.isLoading {
			m.spinner, b.spCmd = m.spinner.Update(msg)
			b.cmds = append(b.cmds, b.spCmd)

			return m, tea.Batch(b.cmds...)
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		m.viewport.Height = msg.Height - 8
		m.viewport.Width = msg.Width
		m.textarea.SetWidth(msg.Width)

		if m.isEditing {
			m.editor, b.editorCmd = m.editor.Update(msg)
			b.cmds = append(b.cmds, b.editorCmd)
		}

		m.textarea, b.tiCmd = m.textarea.Update(msg)
		m.viewport, b.vpCmd = m.viewport.Update(msg)

		b.cmds = append(b.cmds, b.tiCmd, b.vpCmd)
	}

	return m, tea.Batch(b.cmds...)
}
