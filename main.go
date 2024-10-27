package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status int

const (
	about status = iota
	projects
	hobbies
)

var models []tea.Model

const (
	model status = iota
	form
)

var (
	mainStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("241"))
)

type Task struct {
	status      status
	title       string
	description string
}

// Task interface methods
func (t Task) FilterValue() string {
	return t.title
}
func (t Task) Title() string {
	return t.title
}
func (t Task) Description() string {
	return t.description
}

// Main list model struct
type Model struct {
	list list.Model
}

func New() *Model {
	return &Model{}
}

func (m *Model) initList(width, height int) {
	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	m.list.Title = "To do"
	m.list.SetItems([]list.Item{
		Task{status: about, title: "About", description: "SOmething about myself"},
		Task{status: projects, title: "Projects", description: "What Ive been tinkering on"},
		Task{status: hobbies, title: "Hobbies", description: "What I like to do"},
	})
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.initList(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			status := m.list.SelectedItem().(Task).status
			if status == about {
				models[model] = m
				return models[form].Update("This is something about me")
			}
			if status == projects {
				models[model] = m
				return models[form].Update("These are my projects")
			}
			if status == hobbies {
				models[model] = m
				return models[form].Update("These are my hobbies")
			}
			models[model] = m
			return models[form].Update(nil)
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return mainStyle.Render(m.list.View())
}

func main() {
	models = []tea.Model{New(), NewForm()}
	m := models[model]
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
