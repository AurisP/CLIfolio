package main

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	padding   = 2
	maxWidth  = 80
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

	// Input field styles
	inputStyle        = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#87CEEB")) // Blue
	inputHistoryStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#87CEEB"))            // Blue for input history
	outputStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700"))            // Gold for output history
	borderStyle       = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1)
)
