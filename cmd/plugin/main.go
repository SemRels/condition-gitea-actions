package main

import (
	"fmt"
	"io"
	"os"

	plugin "github.com/SemRels/condition-gitea-actions/internal/plugin"
)

const pluginSchemaVersion = 1

func run(getenv func(string) string, stderr io.Writer) int {
	_, _ = fmt.Fprintf(stderr, "plugin_schema_version=%d\n", pluginSchemaVersion)
	c := plugin.NewWithEnv(getenv)
	if err := c.Check(); err != nil {
		fmt.Fprintln(stderr, "condition-gitea-actions:", err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(run(os.Getenv, os.Stderr))
}
