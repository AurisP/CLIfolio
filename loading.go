package main

import (
	"math/rand"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type loadingModel struct {
	progress progress.Model
}

type tickMsg time.Time

func (m loadingModel) Init() tea.Cmd {
	return tickCmd()
}

func (m loadingModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case tickMsg:
		if m.progress.Percent() == 1.0 {
			// Transition to the terminal model when loading is complete
			return NewModel(), nil
		}
		cmd := m.progress.IncrPercent(0.1 + rand.Float64()*(0.8-0.1))
		return m, tea.Batch(tickCmd(), cmd)

	// Handle progress frame updates
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	default:
		return m, nil
	}
}

func (m loadingModel) View() string {
	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + "Loading terminal...\n\n" + // Loading text added here
		pad + m.progress.View() + "\n\n" +
		pad + helpStyle("Press any key to quit")
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
