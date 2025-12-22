package tui

import (
	"bufio"
	"os"

	"github.com/charmbracelet/lipgloss"

	"github.com/thought2code/godev/internal/strconst"
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

func ReadUserInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	return strconst.Empty, scanner.Err()
}
