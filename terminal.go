package main

import (
	"cli/commands"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	input                 textinput.Model
	outputHistory         []string
	displayedHistoryCount int
	revealPosition        int // Tracks how many characters of the current line are revealed
	typingDelay           time.Duration
}

type revealMsg struct{}

func NewModel() *model {
	ti := textinput.New()
	ti.Placeholder = "Type a command..."
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 30

	return &model{
		input:                 ti,
		outputHistory:         []string{"Welcome to the CLI! Type 'help' to see available commands."},
		displayedHistoryCount: 1, // Show the first line
		revealPosition:        0, // Start reveal from the first character
		typingDelay:           5 * time.Millisecond,
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
				m.outputHistory = append(m.outputHistory, "user> "+command)
				response := m.handleCommand(command)
				m.outputHistory = append(m.outputHistory, response)
				m.input.SetValue("")

				// Reset for new output animation
				m.displayedHistoryCount = len(m.outputHistory)
				m.revealPosition = 0
				return m, tea.Tick(m.typingDelay, func(t time.Time) tea.Msg {
					return revealMsg{}
				})
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
		return commands.Help()
	case "about":
		return commands.About()
	case "career":
		return commands.Career()
	case "clear":
		m.outputHistory = nil
		return ""
	case "exit":
		return "Exiting... Goodbye!"
	default:
		return "Unknown command: " + command
	}
}

func (m model) View() string {
	var renderedHistory []string
	visibleCount := m.displayedHistoryCount

	for i, entry := range m.outputHistory {
		if i >= visibleCount {
			break
		}

		if i == visibleCount-1 && m.revealPosition < len(entry) {
			renderedHistory = append(renderedHistory, outputStyle.Render(entry[:m.revealPosition]))
		} else {
			if strings.HasPrefix(entry, "user> ") {
				renderedHistory = append(renderedHistory, inputHistoryStyle.Render(entry))
			} else {
				renderedHistory = append(renderedHistory, outputStyle.Render(entry))
			}
		}
	}

	history := strings.Join(renderedHistory, "\n")
	inputView := inputStyle.Render("user" + m.input.View())
	return borderStyle.Render(history + "\n" + inputView)
}
