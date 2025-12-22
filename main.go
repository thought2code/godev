package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"

	"github.com/thought2code/godev/cmd"
	"github.com/thought2code/godev/internal/strconst"
)

//go:embed template/*
var embedFS embed.FS

var errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(strconst.AnsiColorBrightRed)).Render

func main() {
	cmd.TemplateFS = embedFS
	if err := cmd.Execute(); err != nil {
		fmt.Println(errorStyle(fmt.Sprintf("%s Failed to execute godev: %s", strconst.EmojiFailure, err.Error())))
		os.Exit(1)
	}
}
