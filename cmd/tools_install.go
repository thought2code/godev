package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/thought2code/godev/internal/osutil"
	"github.com/thought2code/godev/internal/strconst"
	"github.com/thought2code/godev/internal/tui"
)

var toolsInstallCmdExample = strings.Trim(`
  godev tools install
  godev tools install golang.org/x/tools/cmd/goimports
`, strconst.Newline)

var toolsInstallCmd = &cobra.Command{
	Use:     "install [tool-package-path]",
	Short:   "Install Go tools",
	Example: toolsInstallCmdExample,
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var toolPkgPath string
		if len(args) > 0 {
			installGoTools(toolPkgPath, strconst.RecommendedGofumptVersion)
		} else {
			fmt.Print(tui.WarnStyle(strconst.EmojiTips + " No tool package path provided. Install recommended tools? (Y/n): "))
			if confirm, err := tui.ReadUserInput(); err != nil {
				fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to read input: %s", strconst.EmojiFailure, err.Error())))
				return
			} else if confirm != "Y" && confirm != "y" {
				fmt.Println(tui.WarnStyle(strconst.EmojiWarning + " godev tools install cancelled"))
				return
			}
			installGoTools(strconst.Gofumpt, strconst.RecommendedGofumptVersion)
			installGoTools(strconst.Goimports, strconst.RecommendedGoimportsVersion)
			installGoTools(strconst.GolangciLint, strconst.RecommendedGolangciLintVersion)
		}
	},
}

func installGoTools(toolName, toolVersion string) {
	fmt.Printf("🔧 Installing %s %s...\n", toolName, toolVersion)
	if err := osutil.RunCommand("go", "install", fmt.Sprintf("%s@%s", toolName, toolVersion)); err != nil {
		fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to install %s %s: %v", strconst.EmojiFailure, toolName, toolVersion, err)))
		return
	}
	fmt.Println(tui.SuccessStyle(fmt.Sprintf("%s Successfully installed %s %s", strconst.EmojiSuccess, toolName, toolVersion)))
}

func init() {
	toolsCmd.AddCommand(toolsInstallCmd)
}
