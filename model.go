// model.go
package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	input         textinput.Model
	commandList   string
	outputHistory []string
}

// NewModel initializes the model with default values and styles.
func NewModel() model {
	ti := textinput.New()
	ti.Placeholder = "Type a command..."
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 30

	// List of available commands
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

// Init initializes the input blinking cursor.
func (m model) Init() tea.Cmd {
	return textinput.Blink
}

// Update processes the key events and handles commands
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			command := strings.TrimSpace(m.input.Value())
			if command != "" {
				m.outputHistory = append(m.outputHistory, "> "+command)
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

// handleCommand processes different commands and returns responses
func (m *model) handleCommand(command string) string {
	switch command {
	case "help":
		return m.commandList
	case "about":
		return "I'm a software developer with a passion for building CLI applications using Go and Bubble Tea!"
	case "projects":
		return "Current projects:\n  - CLI for task management\n  - Portfolio website\n  - Automation scripts for data analysis"
	case "hobbies":
		return "In my free time, I enjoy hiking, reading sci-fi novels, and experimenting with new coding frameworks."
	case "clear":
		m.outputHistory = nil
		return ""
	case "exit":
		return "Exiting... Goodbye!"
	default:
		return "Unknown command: " + command
	}
}

// View renders the UI with separate styles for input and output history
func (m model) View() string {
	var renderedHistory []string

	for _, entry := range m.outputHistory {
		if strings.HasPrefix(entry, "> ") {
			// Render input history in blue
			renderedHistory = append(renderedHistory, inputHistoryStyle.Render(entry))
		} else {
			// Render output responses in gold
			renderedHistory = append(renderedHistory, outputStyle.Render(entry))
		}
	}

	history := strings.Join(renderedHistory, "\n")
	inputView := inputStyle.Render(m.input.View())
	return borderStyle.Render(history + "\n\n" + inputView)
}
