package cmd

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/thought2code/godev/internal/osutil"
)

// global variable to hold the embedded filesystem, initialized in main.go
var TemplateFS embed.FS

const CurrentDir = "."

var initCmd = &cobra.Command{
	Use:   "init [project_name]",
	Short: "Initialize a new Go project from template",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := CurrentDir
		if len(args) > 0 {
			projectName = args[0]
		}

		absPath, _ := filepath.Abs(projectName)

		if projectName == CurrentDir {
			fmt.Println(warningStyle("Warning: Initializing project in current directory may overwrite existing files."))
			fmt.Print(warningStyle("Are you sure to continue? (Y/n): "))

			var confirm string
			fmt.Scan(&confirm)
			if confirm != "Y" && confirm != "y" {
				fmt.Println(warningStyle("godev init cancelled"))
				return
			}

			fmt.Println(warningStyle(fmt.Sprintf("Clearing directory %s", absPath)))
			if err := osutil.ClearDir(absPath); err != nil {
				fmt.Println(errorStyle(fmt.Sprintf("Failed to clear directory: %s, %s", absPath, err.Error())))
				return
			}
		} else {
			if _, err := os.Stat(absPath); err == nil {
				fmt.Println(warningStyle(fmt.Sprintf("Project directory %s already exists", absPath)))
				fmt.Print(warningStyle("Are you sure to overwrite the existing project? (Y/n): "))

				var confirm string
				fmt.Scan(&confirm)
				if confirm != "Y" && confirm != "y" {
					fmt.Println(warningStyle("godev init cancelled"))
					return
				}

				fmt.Println(warningStyle(fmt.Sprintf("Overwriting project directory %s", absPath)))
				if err := os.RemoveAll(absPath); err != nil {
					fmt.Println(errorStyle(fmt.Sprintf("Failed to remove existing project directory: %s", err.Error())))
					return
				}
				if err := os.MkdirAll(absPath, 0o755); err != nil {
					fmt.Println(errorStyle(fmt.Sprintf("Failed to create project directory: %s", err.Error())))
					return
				}
			}
		}

		fmt.Printf("🚀 Creating Go project to: %s\n", absPath)

		files := map[string]string{
			"template/.vscode/launch.json.tpl":   ".vscode/launch.json",
			"template/.vscode/settings.json.tpl": ".vscode/settings.json",
		}

		for src, dest := range files {
			if err := unpack(src, filepath.Join(absPath, dest)); err != nil {
				fmt.Println(errorStyle(fmt.Sprintf("Failed to unpack template file %s: %s", src, err.Error())))
				return
			}
			fmt.Printf("✅ Created file: %s\n", filepath.Join(absPath, dest))
		}

		fmt.Printf("✅ Project initialized successfully: %s\n", absPath)
	},
}

func unpack(src, dest string) error {
	data, err := TemplateFS.ReadFile(src)
	if err != nil {
		return fmt.Errorf("failed to read template file %s: %w", src, err)
	}

	// ensure the directory exists
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return fmt.Errorf("failed to create directory for %s: %w", dest, err)
	}

	if err := os.WriteFile(dest, data, 0o644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", dest, err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)
}
