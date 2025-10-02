package tui

import (
	"log"
	"tasker/internal/gen"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	viewport  viewport.Model
	task      gen.Task
	textarea  textarea.Model
	textStyle lipgloss.Style
	spinner   spinner.Model
	isLoading bool
}

var (
	grayLight  = lipgloss.Color("246")
	grayMedium = lipgloss.Color("240")
	grayDark   = lipgloss.Color("236")
	grayDarker = lipgloss.Color("235")

	softBlue = lipgloss.Color("33")
)

func initialModel() model {
	vp := viewport.New(130, 1)

	ta := textarea.New()
	ta.Focus()
	ta.CharLimit = 200
	ta.SetWidth(vp.Width)
	ta.SetHeight(1)
	ta.ShowLineNumbers = false
	ta.KeyMap.InsertNewline.SetEnabled(false)
	ta.Prompt = "> "

	ta.FocusedStyle.Base = lipgloss.NewStyle().
		Padding(1, 1)

	ta.FocusedStyle.Prompt = lipgloss.NewStyle().
		Foreground(softBlue)

	ta.FocusedStyle.Text = lipgloss.NewStyle().
		Foreground(grayLight)

	ta.FocusedStyle.Placeholder = lipgloss.NewStyle().
		Foreground(grayMedium)

	ta.BlurredStyle.Base = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(grayDark).
		Padding(0, 1)

	ta.BlurredStyle.Prompt = lipgloss.NewStyle().
		Foreground(grayMedium).
		Background(grayDarker)

	ta.BlurredStyle.Text = lipgloss.NewStyle().
		Foreground(lipgloss.Color("244"))

	ta.BlurredStyle.Placeholder = lipgloss.NewStyle().
		Foreground(lipgloss.Color("238"))

	textStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("15")).
		BorderStyle(lipgloss.Border{}).
		Align(lipgloss.Left)

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(softBlue)

	return model{
		textarea:  ta,
		viewport:  vp,
		textStyle: textStyle,
		isLoading: false,
		spinner:   s,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textarea.Blink, m.spinner.Tick)
}

func (m model) View() string {
	var content string

	viewportRender := m.viewport.View()

	inputArea := m.textarea.View()
	if m.isLoading {
		inputArea = lipgloss.NewStyle().Padding(1, 1).Align(lipgloss.Center).Render(
			lipgloss.JoinHorizontal(lipgloss.Center, m.spinner.View(), " Generating response..."),
		)
	}

	content = lipgloss.JoinVertical(lipgloss.Left, inputArea, viewportRender)

	return content
}

func Run() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatalln(err)
	}
}
