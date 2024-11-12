package main

import (
	"cli/commands"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

type model struct {
	input                 textinput.Model
	outputHistory         []string
	displayedHistoryCount int
}

var username string = "Auris@pc"

func NewModel() *model {
	ti := textinput.New()
	ti.Placeholder = "Type a command..."
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 30

	return &model{
		input:                 ti,
		outputHistory:         []string{commands.Welcome()},
		displayedHistoryCount: 1, // Show the first line
	}
}

func (m *model) Init() tea.Cmd {
	// No typing animation needed, so we can return nil
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			command := strings.TrimSpace(m.input.Value())
			if command != "" {
				m.outputHistory = append(m.outputHistory, username+command)
				response := m.handleCommand(command)
				m.outputHistory = append(m.outputHistory, response)
				m.input.SetValue("")

				// Ensure the output is correctly displayed without duplication
				m.displayedHistoryCount = len(m.outputHistory)
			}

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	// Update the input field
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m *model) handleCommand(command string) string {
	switch command {
	case "help":
		return renderMarkdown(commands.Help())
	case "about":
		return renderMarkdown(commands.About())
	case "career":
		return renderMarkdown(commands.Career())
	case "projects":
		return renderMarkdown(commands.Projects())
	case "contacts":
		return renderMarkdown(commands.Contacts())
	case "clear":
		m.outputHistory = nil
		return ""
	case "exit":
		return "Exiting... Goodbye!"
	default:
		return "Unknown command: " + command
	}
}

func renderMarkdown(markdownText string) string {
	renderedText, err := glamour.Render(markdownText, "dark")
	if err != nil {
		log.Println("Error rendering with Glamour:", err)
		return "Error displaying content."
	}
	return strings.TrimSpace(renderedText)
}

func renderEntry(entry string) string {
	if strings.HasPrefix(entry, username) {
		prefix := userStyle.Render(username)
		userInput := textStyle.Render(entry[len(username):]) // Render the user input separately
		return prefix + textStyle.Render("> ") + userInput
	}
	return textStyle.Render(entry)
}

func (m model) View() string {
	var renderedHistory []string
	visibleCount := m.displayedHistoryCount

	for i, entry := range m.outputHistory {
		if i >= visibleCount {
			break
		}

		// Render the entry as normal without reveal animation
		renderedHistory = append(renderedHistory, renderEntry(entry))
	}

	// Join the rendered history and create the input view
	history := strings.Join(renderedHistory, "\n")
	prefix := userStyle.Render(username)
	userInput := textStyle.Render(m.input.View())
	inputView := prefix + userInput

	return borderStyle.Render(history + "\n" + inputView)
}
