package main

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	padding   = 2
	maxWidth  = 80
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626269")).Render

	// Input field styles
	inputStyle        = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#33c5ff")) // Blue
	inputHistoryStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#33c5ff"))            // Blue for input history
	outputStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#c9c7c7"))            // Gold for output history
	borderStyle       = lipgloss.NewStyle()
)
