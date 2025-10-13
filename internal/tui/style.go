package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/lipgloss"
)

var (
	grayLight  = lipgloss.Color("246")
	grayMedium = lipgloss.Color("240")
	grayDarker = lipgloss.Color("235")

	softBlue = lipgloss.Color("33")
)

func styleTextarea(ta *textarea.Model, width int) {
	ta.Focus()
	ta.CharLimit = 200
	ta.SetWidth(width)
	ta.SetHeight(2)
	ta.ShowLineNumbers = false
	ta.KeyMap.InsertNewline.SetEnabled(true)

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
}

func styleSpinner(s *spinner.Model) {
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(softBlue)
}

func styleTips() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(grayMedium).
		Padding(1)
}

func styleTitle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(softBlue)
}
