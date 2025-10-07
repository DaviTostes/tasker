package tui

import (
	"fmt"
	"log"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/kujtimiihoxha/vimtea"
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
		tiCmd     tea.Cmd
		vpCmd     tea.Cmd
		spCmd     tea.Cmd
		editorCmd tea.Cmd
		cmds      []tea.Cmd
	)

	if m.isEditing {
		m.editor, editorCmd = m.editor.Update(msg)
		cmds = append(cmds, editorCmd)
	} else {
		m.textarea, tiCmd = m.textarea.Update(msg)
		m.viewport, vpCmd = m.viewport.Update(msg)
		cmds = append(cmds, tiCmd, vpCmd)
	}

	if m.isLoading {
		m.spinner, spCmd = m.spinner.Update(msg)
		cmds = append(cmds, spCmd)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit

		case tea.KeyEnter:
			if !m.isEditing && !m.isLoading {
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
			if !m.isEditing {
				clipboard.WriteAll(m.task.GetText())
			}

		case tea.KeyCtrlS:
			if m.isEditing {
				if vtEditor, ok := m.editor.(vimtea.Editor); ok {
					m.appContent = vtEditor.GetBuffer().Text()
				} else {
					log.Println("Warning: Editor is not a vimtea.editor, cannot get buffer")
				}

				renderedContent, err := glamour.Render(m.appContent, "dark")
				if err != nil {
					return m, nil
				}

				m.viewport.SetContent(renderedContent)
				m.textarea.Focus()

				m.isEditing = false
				return m, tea.Batch(tiCmd, vpCmd)
			}

		case tea.KeyCtrlE:
			if !m.isEditing {
				m.textarea.Blur()

				m.isEditing = true

				m.editor = vimtea.NewEditor(
					vimtea.WithContent(m.appContent),
					vimtea.WithFileName("task.txt"),
					vimtea.WithFullScreen(),
				)

				m.editor, editorCmd = m.editor.Update(tea.WindowSizeMsg{Width: m.width, Height: m.height})

				cmds = append(cmds, editorCmd)
				cmds = append(cmds, m.editor.Init())

				return m, tea.Batch(cmds...)
			}
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
			m.spinner, spCmd = m.spinner.Update(msg)
			cmds = append(cmds, spCmd)

			return m, tea.Batch(cmds...)
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		m.viewport.Height = msg.Height - 8
		m.viewport.Width = msg.Width
		m.textarea.SetWidth(msg.Width)

		if m.isEditing {
			m.editor, editorCmd = m.editor.Update(msg)
			cmds = append(cmds, editorCmd)
		}

		m.textarea, tiCmd = m.textarea.Update(msg)
		m.viewport, vpCmd = m.viewport.Update(msg)

		cmds = append(cmds, tiCmd, vpCmd)
	}

	return m, tea.Batch(cmds...)
}
