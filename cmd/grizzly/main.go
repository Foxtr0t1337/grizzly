package main

import (
	"os"

	"grizzly/internal/grizzly"
)

func main() {
	cmd := grizzly.NewRootCmd()
	if err := cmd.Execute(); err != nil {
		os.Exit(grizzly.ExitCode(err))
	}
}
