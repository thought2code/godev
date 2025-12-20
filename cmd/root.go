package cmd

import (
	"fmt"
	"runtime"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/thought2code/godev/internal/strconst"
)

var (
	errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color(strconst.AnsiColorBrightRed)).Render
	successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(strconst.AnsiColorBrightGreen)).Render
	warningStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(strconst.AnsiColorBrightYellow)).Render
)

var rootCmd = &cobra.Command{
	Use:   "godev",
	Short: "godev - A modern Go development kit",
	Long:  "godev - A modern Go development kit, helps you to initialize a new Go project from template.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Version = strconst.ProjectVersion

	versionTemplate := fmt.Sprintf(
		strconst.ProjectVersionTemplateFormat,
		time.Now().UTC().Format(strconst.ProjectBuildTimeFormat),
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH)

	rootCmd.SetVersionTemplate(successStyle(versionTemplate))
}
