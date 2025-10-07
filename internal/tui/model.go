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
	textarea   textarea.Model
	spinner    spinner.Model
	editor     tea.Model
	tipsStyle  lipgloss.Style
	task       gen.Task
	appContent string
	isLoading  bool
	isEditing  bool
}

func initialModel() model {
	vp := viewport.New(130, 1)

	ta := textarea.New()
	styleTextarea(&ta, vp.Width)

	s := spinner.New()
	styleSpinner(&s)

	return model{
		textarea:  ta,
		viewport:  vp,
		editor:    vimtea.NewEditor(vimtea.WithFullScreen()),
		spinner:   s,
		tipsStyle: styleTips(),
		isLoading: false,
		isEditing: false,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textarea.Blink, m.spinner.Tick)
}

func (m model) View() string {
	if m.isEditing {
		return m.editor.View()
	}

	viewportRender := m.viewport.View()

	inputArea := m.textarea.View()
	if m.isLoading {
		inputArea = lipgloss.NewStyle().Padding(1, 1, 0).Align(lipgloss.Center).Render(
			lipgloss.JoinHorizontal(lipgloss.Center, m.spinner.View()),
		)
	}

	clipText := m.tipsStyle.Render("Ctrl+y  Clip\nCtrl+e  Edit\nCtrl+s  Exit/Save edit mode\nCtrl+c  Quit")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		inputArea,
		clipText,
		viewportRender,
	)
}

func Run() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatalln(err)
	}
}
