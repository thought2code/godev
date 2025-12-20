package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

const (
	version         = "v0.1.0"
	buildDate       = "2025-12-20"
	versionTemplate = `Version: {{.Name}} {{.Version}} (%s)
Runtime: %s (%s/%s)
Organization: Thought2Code
`
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
	rootCmd.Version = version
	rootCmd.SetVersionTemplate(fmt.Sprintf(
		versionTemplate,
		buildDate,
		runtime.Version(), runtime.GOOS,
		runtime.GOARCH),
	)
}
