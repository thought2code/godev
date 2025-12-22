package tui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	AnsiColorBrightRed    = "1"
	AnsiColorBrightGreen  = "2"
	AnsiColorBrightYellow = "3"
)

var (
	ErrorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color(AnsiColorBrightRed)).Render
	SuccessStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(AnsiColorBrightGreen)).Render
	WarnStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color(AnsiColorBrightYellow)).Render
)
