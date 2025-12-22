package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/thought2code/godev/internal/strconst"
	"github.com/thought2code/godev/internal/tui"
)

var toolsCmd = &cobra.Command{
	Use:     "tools <command>",
	Short:   "Manage Go tools",
	Example: "  godev tools install",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to get help: %s", strconst.EmojiFailure, err.Error())))
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(toolsCmd)
}
