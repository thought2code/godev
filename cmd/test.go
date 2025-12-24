package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/thought2code/godev/internal/osutil"
	"github.com/thought2code/godev/internal/strconst"
	"github.com/thought2code/godev/internal/tui"
)

var testCmdExample = strings.Trim(`
  godev test -u
  godev test -i
`, strconst.Newline)

var (
	runUnitTestFlag        bool
	runIntegrationTestFlag bool
)

const TestCoverageDir = "coverage"

var testCmd = &cobra.Command{
	Use:     "test [-u | -i]",
	Short:   "Run tests for the project",
	Example: testCmdExample,
	Run: func(cmd *cobra.Command, args []string) {
		if runUnitTestFlag {
			runUnitTest()
		}

		if runIntegrationTestFlag {
			runIntegrationTest()
		}
	},
}

func runUnitTest() {
	if err := osutil.RemoveDirIfExist(TestCoverageDir); err != nil {
		fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to remove coverage directory: %s", strconst.EmojiFailure, err.Error())))
		return
	}

	if err := os.MkdirAll(TestCoverageDir, 0o755); err != nil {
		fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to create coverage directory: %s", strconst.EmojiFailure, err.Error())))
		return
	}

	coverprofile := fmt.Sprintf("%s/coverprofile", TestCoverageDir)
	if err := osutil.RunCommand("go", "test", "-v", "-coverprofile", coverprofile, "./..."); err != nil {
		fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to run unit tests: %s", strconst.EmojiFailure, err.Error())))
		return
	}
}

func runIntegrationTest() {
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().BoolVarP(&runUnitTestFlag, "unit", "u", true, "Run unit tests")
	testCmd.Flags().BoolVarP(&runIntegrationTestFlag, "integration", "i", true, "Run integration tests")
}
