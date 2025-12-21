package cmd

import (
	"debug/buildinfo"
	"fmt"
	"os"
	"os/exec"
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
		{
			name:   "Go tools gofumpt",
			result: checkGoToolsVersion("gofumpt"),
		},
		{
			name:   "Go tools goimports",
			result: checkGoToolsVersion("goimports"),
		},
		{
			name:   "Go tools golangci-lint-v2",
			result: checkGoToolsVersion("golangci-lint-v2"),
		},
	}

	for _, c := range checks {
		if c.result.passed {
			fmt.Println(successStyle(fmt.Sprintf("%s %s (%s)", strconst.EmojiSuccess, c.name, c.result.message)))
		} else {
			fmt.Println(errorStyle(fmt.Sprintf("%s %s (%s)", strconst.EmojiFailure, c.name, c.result.message)))
			fmt.Println(warningStyle(fmt.Sprintf("%s Remedy: %s", strconst.EmojiTips, c.result.advice)))
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

func checkGoToolsVersion(toolName string) *checkResult {
	path, err := exec.LookPath(toolName)
	if err != nil {
		return &checkResult{
			passed:  false,
			message: "Not installed",
			advice:  fmt.Sprintf("Run 'godev tidy' to install %s version: %s", toolName, strconst.RecommendedGoimportsVersion),
		}
	}

	version := "unknown"
	info, err := buildinfo.ReadFile(path)
	if err == nil {
		version = info.Main.Version + " built with " + info.GoVersion
	}

	return &checkResult{
		passed:  true,
		message: fmt.Sprintf("installed version: %s", version),
		advice:  strconst.Empty,
	}
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
