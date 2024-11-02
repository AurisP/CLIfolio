package main

import (
	"strings"

	"cli/commands"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	input         textinput.Model
	commandList   string
	outputHistory []string
}

func NewModel() model {
	ti := textinput.New()
	ti.Placeholder = "Type a command..."
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 30

	commandList := `
Available commands:
  - help      Show available commands
  - about     Display information about me
  - projects  List my projects
  - hobbies   Describe my hobbies
  - clear     Clear the screen
  - exit      Exit the application
`

	return model{
		input:         ti,
		commandList:   commandList,
		outputHistory: []string{"Welcome to the CLI! Type 'help' to see available commands."},
	}
}

// CLI Update function
func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			command := strings.TrimSpace(m.input.Value())
			if command != "" {
				m.outputHistory = append(m.outputHistory, "user> "+command)
				response := m.handleCommand(command)
				m.outputHistory = append(m.outputHistory, response)
				m.input.SetValue("") // Clear the input field
			}

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m *model) handleCommand(command string) string {
	switch command {
	case "help":
		return commands.Help()
	case "about":
		return commands.About()
	case "clear":
		m.outputHistory = nil
		return ""
	case "exit":
		return "Exiting... Goodbye!"
	default:
		return "Unknown command: " + command
	}
}

// ... [rest of your code]

// CLI View function
func (m model) View() string {
	var renderedHistory []string

	for _, entry := range m.outputHistory {
		if strings.HasPrefix(entry, "user> ") {
			renderedHistory = append(renderedHistory, inputHistoryStyle.Render(entry))
		} else {
			renderedHistory = append(renderedHistory, outputStyle.Render(entry))
		}
	}

	history := strings.Join(renderedHistory, "\n")
	inputView := inputStyle.Render("user" + m.input.View())
	return borderStyle.Render(history + "\n\n" + inputView)
}
