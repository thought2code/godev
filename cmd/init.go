package cmd

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/thought2code/godev/internal/osutil"
	"github.com/thought2code/godev/internal/strconst"
	"github.com/thought2code/godev/internal/tui"
)

// global variable to hold the embedded filesystem, initialized in main.go
var TemplateFS embed.FS

var initCmdExample = strings.Trim(`
  godev init
  godev init myproject
`, strconst.Newline)

const CurrentDir = "."

var initCmd = &cobra.Command{
	Use:     "init [project-name]",
	Short:   "Initialize a new Go project from template",
	Example: initCmdExample,
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := CurrentDir
		if len(args) > 0 {
			projectName = args[0]
		}
		absPath, _ := filepath.Abs(projectName)

		if !initInDir(absPath) {
			return
		}

		gitRepo := userInputGitRepo()
		if gitRepo == strconst.Empty {
			gitRepo = filepath.Base(absPath)
		}

		fmt.Printf("%s Creating Go project to: %s\n", strconst.EmojiRocket, absPath)
		if !unpackTemplatesAndReplacePlaceholders(absPath, gitRepo) {
			return
		}
		fmt.Printf("%s Project initialized successfully: %s\n", strconst.EmojiSuccess, absPath)
	},
}

func initInDir(dirAbsPath string) (kontinue bool) {
	exist, err := osutil.CheckExist(dirAbsPath)
	if err != nil {
		fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to check directory: %s", strconst.EmojiFailure, err.Error())))
		return false
	}

	if exist {
		empty, err := osutil.CheckDirEmpty(dirAbsPath)
		if err != nil {
			fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to check if directory is empty: %s", strconst.EmojiFailure, err.Error())))
			return false
		}

		if empty {
			return true
		}

		fmt.Println(tui.WarnStyle(fmt.Sprintf("%s Project directory %s is not empty", strconst.EmojiWarning, dirAbsPath)))
		fmt.Print(tui.WarnStyle(strconst.EmojiQuestion + " Are you sure to initialize the project in the specified directory? (Y/n): "))

		if confirm, err := tui.ReadUserInput(); err != nil {
			fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to read input: %s", strconst.EmojiFailure, err.Error())))
			return false
		} else if confirm != "Y" && confirm != "y" {
			fmt.Println(tui.WarnStyle(strconst.EmojiWarning + " godev init cancelled"))
			return false
		}
		return true
	} else {
		if err := os.MkdirAll(dirAbsPath, 0o755); err != nil {
			fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to create project directory: %s", strconst.EmojiFailure, err.Error())))
			return false
		}
		return true
	}
}

func userInputGitRepo() string {
	fmt.Printf("%s Git repository (optional, e.g. github.com/thought2code/godev, press Enter to skip): ", strconst.EmojiQuestion)
	if input, err := tui.ReadUserInput(); err != nil {
		fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to read input: %s", strconst.EmojiFailure, err.Error())))
		return strconst.Empty
	} else {
		return input
	}
}

func unpackTemplatesAndReplacePlaceholders(dirAbsPath, gitRepo string) (success bool) {
	files := map[string]string{
		"template/.vscode/extensions.json.tpl": ".vscode/extensions.json",
		"template/.vscode/launch.json.tpl":     ".vscode/launch.json",
		"template/.vscode/settings.json.tpl":   ".vscode/settings.json",
		"template/.gitignore.tpl":              ".gitignore",
		"template/.golangci.yml.tpl":           ".golangci.yml",
		"template/go.mod.tpl":                  "go.mod",
	}

	replacements := map[string]string{
		"{{.ProjectName}}":     filepath.Base(dirAbsPath),
		"{{.LatestGoVersion}}": fetchLatestGoVersion(),
		"{{.GitRepo}}":         strings.TrimPrefix(gitRepo, "https://"),
	}

	for src, dest := range files {
		if err := unpack(src, filepath.Join(dirAbsPath, dest), replacements); err != nil {
			fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to unpack template file %s: %s", strconst.EmojiFailure, src, err.Error())))
			return false
		}
		fmt.Printf("%s Created file: %s\n", strconst.EmojiSuccess, filepath.Join(dirAbsPath, dest))
	}
	return true
}

func unpack(src, dest string, replacements map[string]string) error {
	bytes, err := TemplateFS.ReadFile(src)
	if err != nil {
		return fmt.Errorf("failed to read template file %s: %w", src, err)
	}

	// ensure the directory exists
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return fmt.Errorf("failed to create directory for %s: %w", dest, err)
	}

	// apply replacements
	for key, value := range replacements {
		bytes = []byte(strings.ReplaceAll(string(bytes), key, value))
	}

	if err := os.WriteFile(dest, bytes, 0o644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", dest, err)
	}

	return nil
}

func fetchLatestGoVersion() string {
	// try to fetch from go.dev first
	resp, err := http.Get("https://go.dev/VERSION?m=text")
	if err != nil {
		// if go.dev fails, try golang.org as fallback
		resp, err = http.Get("https://golang.org/VERSION?m=text")
		if err != nil {
			fmt.Println(tui.WarnStyle(fmt.Sprintf("%s Failed to fetch latest Go version: %s", strconst.EmojiWarning, err.Error())))
			fmt.Println(tui.WarnStyle(fmt.Sprintf("%s Falling back to latest Go version (%s) at the time of godev release", strconst.EmojiWarning, strconst.LatestGoVersionFallback)))
			return strconst.LatestGoVersionFallback
		}
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to close response body: %s", strconst.EmojiFailure, closeErr.Error())))
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to read response body: %s", strconst.EmojiFailure, err.Error())))
		fmt.Println(tui.WarnStyle(fmt.Sprintf("%s Falling back to latest Go version (%s) at the time of godev release", strconst.EmojiWarning, strconst.LatestGoVersionFallback)))
		return strconst.LatestGoVersionFallback
	}

	latestGoVersion := strings.Split(strings.TrimSpace(string(body)), strconst.Newline)[0]
	return strings.TrimPrefix(latestGoVersion, "go")
}

func init() {
	rootCmd.AddCommand(initCmd)
}
