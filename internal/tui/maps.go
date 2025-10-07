package tui

import (
	"log"

	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/kujtimiihoxha/vimtea"
)

type KeyMapFunc func() (tea.Model, tea.Cmd)

func GetKeyMap(m model, b batch) map[tea.KeyType]KeyMapFunc {
	return map[tea.KeyType]KeyMapFunc{
		tea.KeyCtrlC: func() (tea.Model, tea.Cmd) {
			return m, tea.Quit
		},

		tea.KeyEnter: func() (tea.Model, tea.Cmd) {
			if !m.isEditing && !m.isLoading {
				userMessage := m.textarea.Value()
				if userMessage == "" {
					return m, tea.Batch(b.cmds...)
				}

				m.isLoading = true
				m.textarea.Reset()
				m.viewport.GotoBottom()

				return m, tea.Batch(performGeneration(userMessage), m.spinner.Tick)
			}

			return m, tea.Batch(b.cmds...)
		},

		tea.KeyCtrlY: func() (tea.Model, tea.Cmd) {
			if !m.isEditing {
				clipboard.WriteAll(m.appContent)
			}

			return m, tea.Batch(b.cmds...)
		},

		tea.KeyCtrlS: func() (tea.Model, tea.Cmd) {
			if m.isEditing {
				if vtEditor, ok := m.editor.(vimtea.Editor); ok {
					m.appContent = vtEditor.GetBuffer().Text()
				} else {
					log.Println("Warning: Editor is not a vimtea.editor, cannot get buffer")
				}

				renderedContent, err := glamour.Render(m.appContent, "dark")
				if err != nil {
					return m, tea.Batch(b.cmds...)
				}

				m.viewport.SetContent(renderedContent)
				m.textarea.Focus()

				m.isEditing = false
				return m, tea.Batch(b.tiCmd, b.vpCmd)
			}

			return m, tea.Batch(b.cmds...)
		},

		tea.KeyCtrlE: func() (tea.Model, tea.Cmd) {
			if !m.isEditing {
				m.textarea.Blur()

				m.isEditing = true

				m.editor = vimtea.NewEditor(
					vimtea.WithContent(m.appContent),
					vimtea.WithFileName("task.txt"),
					vimtea.WithFullScreen(),
				)

				m.editor, b.editorCmd = m.editor.Update(tea.WindowSizeMsg{Width: m.width, Height: m.height})

				b.cmds = append(b.cmds, b.editorCmd)
				b.cmds = append(b.cmds, m.editor.Init())

				return m, tea.Batch(b.cmds...)
			}

			return m, tea.Batch(b.cmds...)
		},
	}
}
