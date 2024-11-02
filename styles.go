// styles.go
package main

import "github.com/charmbracelet/lipgloss"

// Define styles for input, output, input history, and border
var (
	inputStyle        = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#87CEEB")) // Blue
	inputHistoryStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#87CEEB"))            // Blue for input history
	outputStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700"))            // Gold for output history
	borderStyle       = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1)
)
