package main

import (
	"os"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Start with progress bar model
	m := loadingModel{
		progress: progress.New(progress.WithScaledGradient("#f0f2f2", "#08b9ff")),
	}

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		os.Exit(1)
	}
}
