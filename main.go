package main

import (
	"fmt"
	"io"
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
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	// Assert that listItem is of type item or Task
	i, ok := listItem.(Task)
	if !ok {
		return
	}

	// Format the item as a string
	var str string
	if index == m.Index() {
		// Selected item style
		str = selectedItemStyle.Render(fmt.Sprintf("> %s", i.Title()))
	} else {
		// Regular item style
		str = itemStyle.Render(fmt.Sprintf("%d. %s", index+1, i.Title()))
	}

	// Output the styled item string
	fmt.Fprint(w, str)
}

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
	m.list = list.New([]list.Item{}, itemDelegate{}, width, height)
	m.list.Styles.Title = titleStyle
	m.list.Styles.PaginationStyle = paginationStyle
	m.list.Styles.HelpStyle = helpStyle
	m.list.SetShowStatusBar(false)
	m.list.SetFilteringEnabled(false)

	m.list.Title = "To do"
	m.list.SetItems([]list.Item{
		Task{status: about, title: "About", description: "Something about myself"},
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

	return m.list.View()
}

func main() {
	models = []tea.Model{New(), NewForm()}
	m := models[model]
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
