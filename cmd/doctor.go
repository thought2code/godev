package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/thought2code/godev/internal/strconst"
	"golang.org/x/mod/modfile"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Diagnose the health of the development environment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔍 Diagnosing the development environment...")
		runDoctor()
	},
}

func runDoctor() {
	type check struct {
		name   string
		result *checkResult
	}

	checks := []check{
		{
			name:   "Go module file",
			result: checkGoModuleFile(),
		},
		{
			name:   "Go version",
			result: checkGoVersion(),
		},
	}

	for _, c := range checks {
		if c.result.passed {
			fmt.Println(successStyle(fmt.Sprintf("%s %s (%s)", strconst.EmojiSuccess, c.name, c.result.message)))
		} else {
			fmt.Println(errorStyle(fmt.Sprintf("%s %s (%s)", strconst.EmojiFailure, c.name, c.result.message)))
			fmt.Println(warningStyle(fmt.Sprintf("%s Tips: %s", strconst.EmojiTips, c.result.advice)))
		}
	}
}

type checkResult struct {
	passed  bool
	message string
	advice  string
}

func checkGoModuleFile() *checkResult {
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		return &checkResult{
			passed:  false,
			message: "The go.mod file does not exist",
			advice:  "Create the go.mod file using 'go mod init'",
		}
	}

	return &checkResult{
		passed:  true,
		message: "The go.mod file exists",
		advice:  strconst.Empty,
	}
}

func checkGoVersion() *checkResult {
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		return &checkResult{
			passed:  false,
			message: "No go.mod file found, unable to check required Go version",
			advice:  "Create the go.mod file using 'go mod init'",
		}
	}

	data, err := os.ReadFile("go.mod")
	if err != nil {
		return &checkResult{
			passed:  false,
			message: err.Error(),
			advice:  "Check if the go.mod file exists and readable",
		}
	}

	mod, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return &checkResult{
			passed:  false,
			message: err.Error(),
			advice:  "Check if the go.mod file is a valid Go module file",
		}
	}

	if mod.Go == nil {
		return &checkResult{
			passed:  false,
			message: "The Go version directive is missing in the go.mod file",
			advice:  "Add the Go version directive to the go.mod file",
		}
	}

	requiredGoVersion := "go" + mod.Go.Version
	installedGoVersion := runtime.Version()
	passed := installedGoVersion >= requiredGoVersion
	advice := strconst.Empty
	if !passed {
		advice = fmt.Sprintf("Upgrade Go to version %s or higher, download from https://golang.org/dl/", requiredGoVersion)
	}
	return &checkResult{
		passed:  passed,
		message: fmt.Sprintf("installed: %s, required: %s", installedGoVersion, requiredGoVersion),
		advice:  advice,
	}
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
