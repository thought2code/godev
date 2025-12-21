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
	"github.com/thought2code/godev/internal/strconst"
)

// global variable to hold the embedded filesystem, initialized in main.go
var TemplateFS embed.FS

var initCmd = &cobra.Command{
	Use:     "init <project-name>",
	Short:   "Initialize a new Go project from template",
	Example: "  godev init myproject",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		absPath, _ := filepath.Abs(projectName)

		if _, err := os.Stat(absPath); err == nil {
			fmt.Println(warningStyle(fmt.Sprintf("%s Project directory %s already exists", strconst.EmojiWarning, absPath)))
			fmt.Print(warningStyle(strconst.EmojiQuestion + " Are you sure to overwrite the existing project? (Y/n): "))

			var confirm string
			fmt.Scan(&confirm)
			if confirm != "Y" && confirm != "y" {
				fmt.Println(warningStyle(strconst.EmojiWarning + " godev init cancelled"))
				return
			}

			fmt.Println(warningStyle(fmt.Sprintf("%s Overwriting project directory %s", strconst.EmojiWarning, absPath)))
			if err := os.RemoveAll(absPath); err != nil {
				fmt.Println(errorStyle(fmt.Sprintf("%s Failed to remove existing project directory: %s", strconst.EmojiFailure, err.Error())))
				return
			}
			if err := os.MkdirAll(absPath, 0o755); err != nil {
				fmt.Println(errorStyle(fmt.Sprintf("%s Failed to create project directory: %s", strconst.EmojiFailure, err.Error())))
				return
			}
		}

		fmt.Printf("%s Creating Go project to: %s\n", strconst.EmojiRocket, absPath)

		files := map[string]string{
			"template/.vscode/extensions.json.tpl": ".vscode/extensions.json",
			"template/.vscode/launch.json.tpl":     ".vscode/launch.json",
			"template/.vscode/settings.json.tpl":   ".vscode/settings.json",
			"template/go.mod.tpl":                  "go.mod",
		}

		replacements := map[string]string{
			"{{.ProjectName}}":     filepath.Base(absPath),
			"{{.LatestGoVersion}}": fetchLatestGoVersion(),
		}

		for src, dest := range files {
			if err := unpack(src, filepath.Join(absPath, dest), replacements); err != nil {
				fmt.Println(errorStyle(fmt.Sprintf("%s Failed to unpack template file %s: %s", strconst.EmojiFailure, src, err.Error())))
				return
			}
			fmt.Printf("%s Created file: %s\n", strconst.EmojiSuccess, filepath.Join(absPath, dest))
		}

		fmt.Printf("%s Project initialized successfully: %s\n", strconst.EmojiSuccess, absPath)
	},
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
			fmt.Println(warningStyle(fmt.Sprintf("%s Failed to fetch latest Go version: %s", strconst.EmojiWarning, err.Error())))
			fmt.Println(warningStyle(fmt.Sprintf("%s Falling back to latest Go version (%s) at the time of godev release", strconst.EmojiWarning, strconst.LatestGoVersionFallback)))
			return strconst.LatestGoVersionFallback
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(errorStyle(fmt.Sprintf("%s Failed to read response body: %s", strconst.EmojiFailure, err.Error())))
		fmt.Println(warningStyle(fmt.Sprintf("%s Falling back to latest Go version (%s) at the time of godev release", strconst.EmojiWarning, strconst.LatestGoVersionFallback)))
		return strconst.LatestGoVersionFallback
	}

	latestGoVersion := strings.Split(strings.TrimSpace(string(body)), strconst.Newline)[0]
	return strings.TrimPrefix(latestGoVersion, "go")
}

func init() {
	rootCmd.AddCommand(initCmd)
}
