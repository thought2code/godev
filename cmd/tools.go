package cmd

import (
	"github.com/spf13/cobra"
)

var toolsCmd = &cobra.Command{
	Use:     "tools <command>",
	Short:   "Manage Go tools",
	Example: "  godev tools install",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(toolsCmd)
}
