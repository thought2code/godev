package cmd

import (
	"fmt"
	"runtime"
	"strings"
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
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			fmt.Println(errorStyle(fmt.Sprintf("%s Failed to get help: %s", strconst.EmojiFailure, err.Error())))
			return
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Version = strconst.ProjectVersion

	versionTemplateFormat := strings.Trim(strconst.ProjectVersionTemplateFormat, strconst.Newline)
	buildTime := time.Now().UTC().Format(strconst.ProjectBuildTimeFormat)
	versionTemplate := fmt.Sprintf(
		versionTemplateFormat,
		buildTime,
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH)

	rootCmd.SetVersionTemplate(successStyle(versionTemplate))
}
