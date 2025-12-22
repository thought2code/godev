package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thought2code/godev/internal/osutil"
	"github.com/thought2code/godev/internal/strconst"
)

var example = strings.Trim(`
  godev tools install
  godev tools install golang.org/x/tools/cmd/goimports
`, strconst.Newline)

var toolsInstallCmd = &cobra.Command{
	Use:     "install [tool-package-path]",
	Short:   "Install Go tools",
	Example: example,
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var toolPkgPath string
		if len(args) > 0 {
			installGoTools(toolPkgPath, strconst.RecommendedGofumptVersion)
		} else {
			fmt.Print(warningStyle(strconst.EmojiTips + " No tool package path provided. Install recommended tools? (Y/n): "))
			var confirm string
			fmt.Scan(&confirm)
			if confirm != "Y" && confirm != "y" {
				fmt.Println(warningStyle(strconst.EmojiWarning + " godev tools install cancelled"))
				return
			}
			installGoTools(strconst.Gofumpt, strconst.RecommendedGofumptVersion)
			installGoTools(strconst.Goimports, strconst.RecommendedGoimportsVersion)
			installGoTools(strconst.GolangciLint, strconst.RecommendedGolangciLintVersion)
		}
	},
}

func installGoTools(toolName string, toolVersion string) {
	fmt.Printf("🔧 Installing %s %s...\n", toolName, toolVersion)
	if err := osutil.RunCommand("go", "install", fmt.Sprintf("%s@%s", toolName, toolVersion)); err != nil {
		fmt.Println(errorStyle(fmt.Sprintf("%s Failed to install %s %s: %v", strconst.EmojiFailure, toolName, toolVersion, err)))
		return
	}
	fmt.Println(successStyle(fmt.Sprintf("%s Successfully installed %s %s", strconst.EmojiSuccess, toolName, toolVersion)))
}

func init() {
	toolsCmd.AddCommand(toolsInstallCmd)
}
