package main

import (
	"os"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Start with progress bar model
	m := loadingModel{
		progress: progress.New(progress.WithScaledGradient("#FF7CCB", "#FDFF8C")),
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		os.Exit(1)
	}
}
