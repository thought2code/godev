package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"

	"github.com/thought2code/godev/internal/osutil"
	"github.com/thought2code/godev/internal/strconst"
	"github.com/thought2code/godev/internal/tui"
)

var unitTestCmdExample = strings.Trim(`
  godev test unit
  godev test unit -v
  godev test unit -c
  godev test unit -v -c
  godev test unit --html
  godev test unit -v --html
`, strconst.NewLine)

var (
	verboseFlag    bool
	coverageFlag   bool
	htmlReportFlag bool
)

var unitTestCmd = &cobra.Command{
	Use:     "unit [-v] [-c] [--html]",
	Short:   "Run unit tests for the project",
	Example: unitTestCmdExample,
	Run: func(cmd *cobra.Command, args []string) {
		testCoverageDir := "coverage"

		if err := osutil.RemoveDirIfExist(testCoverageDir); err != nil {
			fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to remove coverage directory: %s", strconst.EmojiFailure, err.Error())))
			return
		}

		if err := os.MkdirAll(testCoverageDir, 0o755); err != nil {
			fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to create coverage directory: %s", strconst.EmojiFailure, err.Error())))
			return
		}

		coverprofile := filepath.Join(testCoverageDir, "coverprofile")

		testArgs := []string{"test"}
		if verboseFlag {
			testArgs = append(testArgs, "-v")
		}
		if coverageFlag || htmlReportFlag {
			testArgs = append(testArgs, "-coverprofile", coverprofile)
		}
		testArgs = append(testArgs, "./...")

		if err := osutil.RunCommand("go", testArgs...); err != nil {
			fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to run unit tests: %s", strconst.EmojiFailure, err.Error())))
			return
		}

		if htmlReportFlag {
			html := filepath.Join(testCoverageDir, "cover.html")
			if err := osutil.RunCommand("go", "tool", "cover", "-html", coverprofile, "-o", html); err != nil {
				fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to generate HTML coverage report: %s", strconst.EmojiFailure, err.Error())))
				return
			}
			if runtime.GOOS == "windows" {
				if err := osutil.RunCommand("cmd", "/c", "start", html); err != nil {
					fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to open HTML coverage report: %s", strconst.EmojiFailure, err.Error())))
					return
				}
			}
			fmt.Println(tui.SuccessStyle(fmt.Sprintf("%s HTML coverage report generated: %s", strconst.EmojiSuccess, html)))
		}
	},
}

func init() {
	testCmd.AddCommand(unitTestCmd)
	unitTestCmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "Enable verbose output")
	unitTestCmd.Flags().BoolVarP(&coverageFlag, "cover", "c", false, "Enable code coverage")
	unitTestCmd.Flags().BoolVar(&htmlReportFlag, "html", false, "Generate and open HTML coverage report")
}
