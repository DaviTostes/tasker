package model

import (
	"fmt"
	"strings"
	"tasker/internal/gen"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

type model struct {
	viewport    viewport.Model
	messages    []string
	textarea    textarea.Model
	senderStyle lipgloss.Style
	botStyle    lipgloss.Style
	textStyle   lipgloss.Style
}

func InitialModel() model {
	ta := textarea.New()
	ta.Placeholder = "Enter your task here..."
	ta.Focus()

	ta.Prompt = lipgloss.NewStyle().Foreground(lipgloss.Color("5")).Render("| ")
	ta.CharLimit = 200

	ta.SetWidth(50)
	ta.SetHeight(1)

	ta.ShowLineNumbers = false

	vp := viewport.New(130, 5)

	initialMessage := []string{"Welcome to Tasker!"}
	vp.SetContent(wordwrap.String(strings.Join(initialMessage, "\n"), 130))

	ta.KeyMap.InsertNewline.SetEnabled(false)

	senderStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("5")).Align(lipgloss.Left)
	botStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("6")).Align(lipgloss.Left)
	textStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("15")).Align(lipgloss.Left)

	return model{
		textarea:    ta,
		messages:    initialMessage,
		viewport:    vp,
		senderStyle: senderStyle,
		botStyle:    botStyle,
		textStyle:   textStyle,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			userMessage := m.textarea.Value()
			m.messages = append(m.messages, m.senderStyle.Render("| "+m.textStyle.Render(userMessage)))

			botResponse, err := gen.GenerateTask(userMessage)
			if err != nil {
				fmt.Println(err)
				return m, tea.Quit
			}

			m.messages = append(m.messages, m.botStyle.Render("| "+m.textStyle.Render(botResponse)))

			m.viewport.SetContent(wordwrap.String(strings.Join(m.messages, "\n\n"), 130))

			m.textarea.Reset()
			m.viewport.GotoBottom()
		}

	case tea.WindowSizeMsg:
		m.viewport.Height = msg.Height - 3
		m.viewport.Width = msg.Width
		m.textarea.SetWidth(msg.Width)

		m.viewport.SetContent(wordwrap.String(strings.Join(m.messages, "\n\n"), m.viewport.Width))
	}

	return m, tea.Batch(tiCmd, vpCmd)
}

func (m model) View() string {
	return fmt.Sprintf("%s\n\n%s", m.viewport.View(), m.textarea.View())
}
