package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/thought2code/godev/cmd"
	"github.com/thought2code/godev/internal/strconst"
	"github.com/thought2code/godev/internal/tui"
)

//go:embed template/*
var embedFS embed.FS

func main() {
	cmd.TemplateFS = embedFS
	if err := cmd.Execute(); err != nil {
		fmt.Println(tui.ErrorStyle(fmt.Sprintf("%s Failed to execute godev: %s", strconst.EmojiFailure, err.Error())))
		os.Exit(1)
	}
}
