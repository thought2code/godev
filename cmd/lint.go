package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/thought2code/godev/internal/osutil"
	"github.com/thought2code/godev/internal/strconst"
	"github.com/thought2code/godev/internal/tui"
)

var lintCmd = &cobra.Command{
	Use:     "lint",
	Short:   "Run linters on the codebase",
	Example: "  godev lint",
	PreRun: func(cmd *cobra.Command, args []string) {
		runDoctor()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := osutil.RunCommand("goimports", "-w", "."); err != nil {
			fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to run 'goimports -w .': %v", strconst.EmojiFailure, err)))
			return
		}
		if err := osutil.RunCommand("gofumpt", "-w", "."); err != nil {
			fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to run 'gofumpt -w .': %v", strconst.EmojiFailure, err)))
			return
		}
		if err := osutil.RunCommand("golangci-lint", "run", "./..."); err != nil {
			fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to run 'golangci-lint run ./...': %v", strconst.EmojiFailure, err)))
			return
		}
		if err := osutil.RunCommand("go", "mod", "tidy"); err != nil {
			fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to run 'go mod tidy': %v", strconst.EmojiFailure, err)))
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(lintCmd)
}
