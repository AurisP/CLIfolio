package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true).
			Padding(0, 1)

	bodyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("244")).
			Padding(1, 1)
)

type Form struct{}

func NewForm() *Form {
	return &Form{}
}

func (m Form) Init() tea.Cmd {
	return nil
}

func (m Form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Form) View() string {

	title := "About myself"
	body := "Something about myself"

	return titleStyle.Render(title) + "\n\n" + bodyStyle.Render(body)
}
