package main

import (
	"embed"

	"github.com/thought2code/godev/cmd"
)

//go:embed template/*
var embedFS embed.FS

func main() {
	cmd.TemplateFS = embedFS
	cmd.Execute()
}
