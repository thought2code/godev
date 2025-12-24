package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/thought2code/godev/internal/strconst"
	"github.com/thought2code/godev/internal/tui"
)

var testCmdExample = strings.Trim(`
  godev test unit
  godev test integ
`, strconst.Newline)

var testCmd = &cobra.Command{
	Use:     "test",
	Short:   "Run tests for the project",
	Example: testCmdExample,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to get help: %s", strconst.EmojiFailure, err.Error())))
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
