package main

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	padding   = 2
	maxWidth  = 80
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626269")).Render

	// Input field styles
	userStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffbb00")) // Blue
	textStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#c9c7c7"))
	borderStyle = lipgloss.NewStyle()
)
