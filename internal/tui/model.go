package tui

import (
	"log"
	"tasker/internal/gen"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kujtimiihoxha/vimtea"
)

type model struct {
	width      int
	height     int
	viewport   viewport.Model
	task       gen.Task
	textarea   textarea.Model
	spinner    spinner.Model
	isLoading  bool
	editor     tea.Model
	isEditing  bool
	appContent string
	textStyle  lipgloss.Style
	tipsStyle  lipgloss.Style
}

var (
	grayLight  = lipgloss.Color("246")
	grayMedium = lipgloss.Color("240")
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
		Padding(1, 1, 0, 1)

	ta.FocusedStyle.Prompt = lipgloss.NewStyle().
		Foreground(softBlue)

	ta.FocusedStyle.Text = lipgloss.NewStyle().
		Foreground(grayLight)

	ta.FocusedStyle.Placeholder = lipgloss.NewStyle().
		Foreground(grayMedium)

	ta.BlurredStyle.Prompt = lipgloss.NewStyle().
		Foreground(grayMedium).
		Background(grayDarker)

	ta.BlurredStyle.Text = lipgloss.NewStyle().
		Foreground(lipgloss.Color("244"))

	ta.BlurredStyle.Placeholder = lipgloss.NewStyle().
		Foreground(lipgloss.Color("238"))

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(softBlue)

	textStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("15")).
		BorderStyle(lipgloss.Border{}).
		Align(lipgloss.Left)

	tipsStyle := lipgloss.NewStyle().
		Foreground(grayMedium).
		Padding(1)

	return model{
		textarea:  ta,
		viewport:  vp,
		isLoading: false,
		spinner:   s,
		editor:    vimtea.NewEditor(vimtea.WithFullScreen()),
		isEditing: false,
		textStyle: textStyle,
		tipsStyle: tipsStyle,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textarea.Blink, m.spinner.Tick)
}

func (m model) View() string {
	var content string

	if m.isEditing {
		editorView := m.editor.View()
		content = lipgloss.JoinVertical(lipgloss.Left,
			m.tipsStyle.Render("Ctrl+S to save/exit"),
			lipgloss.NewStyle().Render(editorView),
		)

		return content
	}

	viewportRender := m.viewport.View()

	inputArea := m.textarea.View()
	if m.isLoading {
		inputArea = lipgloss.NewStyle().Padding(1, 1, 0).Align(lipgloss.Center).Render(
			lipgloss.JoinHorizontal(lipgloss.Center, m.spinner.View()),
		)
	}

	clipText := m.tipsStyle.Render("Ctrl+y  Clip\nCtrl+e  Edit\nCtrl+s  Exit/Save edit mode\nCtrl+c  Quit")

	content = lipgloss.JoinVertical(
		lipgloss.Left,
		inputArea,
		clipText,
		viewportRender,
	)

	return content
}

func Run() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatalln(err)
	}
}
