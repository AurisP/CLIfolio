package main

import (
	"cli/commands"
	"log"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

type model struct {
	input                 textinput.Model
	outputHistory         []string
	displayedHistoryCount int
	revealPosition        int // Tracks how many characters of the current line are revealed
	typingDelay           time.Duration
	displaySlow           bool
}

type revealMsg struct{}

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
		revealPosition:        0, // Start reveal from the first character
		typingDelay:           2 * time.Millisecond,
	}
}

func (m model) Init() tea.Cmd {
	// Start the reveal animation immediately on application load
	return tea.Tick(m.typingDelay, func(t time.Time) tea.Msg {
		return revealMsg{}
	})
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

				// Reset for new output animation
				m.displayedHistoryCount = len(m.outputHistory)
				m.revealPosition = 0

				if m.displaySlow {
					return m, tea.Tick(m.typingDelay, func(t time.Time) tea.Msg {
						return revealMsg{}
					})
				} else {
					m.revealPosition = len(m.outputHistory[m.displayedHistoryCount-1])
				}

			}

		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case revealMsg:
		// Increment reveal position and continue until line is fully displayed
		if m.revealPosition < len(m.outputHistory[m.displayedHistoryCount-1]) {
			m.revealPosition++
			return m, tea.Tick(m.typingDelay, func(t time.Time) tea.Msg {
				return revealMsg{}
			})
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m *model) handleCommand(command string) string {
	switch command {
	case "help":
		m.displaySlow = false
		return renderMarkdown(commands.Help())
	case "about":
		m.displaySlow = false
		return renderMarkdown(commands.About())
	case "career":
		m.displaySlow = false
		return renderMarkdown(commands.Career())
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

		if i == visibleCount-1 && m.revealPosition < len(entry) {
			// Render the partially revealed entry with a style
			renderedHistory = append(renderedHistory, textStyle.Render(entry[:m.revealPosition]))
		} else {
			// Use the helper function to render the entry
			renderedHistory = append(renderedHistory, renderEntry(entry))
		}
	}

	// Join the rendered history and create the input view
	history := strings.Join(renderedHistory, "\n")
	prefix := userStyle.Render(username)
	userInput := textStyle.Render(m.input.View())
	inputView := prefix + userInput

	return borderStyle.Render(history + "\n" + inputView)
}
