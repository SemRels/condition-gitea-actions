package plugin

import (
	"fmt"
	"os"
	"strings"
)

type Condition struct {
	env func(string) string
}

func New() *Condition { return &Condition{env: os.Getenv} }

func NewWithEnv(env func(string) string) *Condition { return &Condition{env: env} }

func (c *Condition) Check() error {
	var errs []string

	if c.env("GITEA_ACTIONS") != "true" {
		errs = append(errs, "GITEA_ACTIONS is not set to \"true\"; this plugin requires a Gitea Actions environment")
	}

	if c.env("GITEA_TOKEN") == "" && c.env("CI_JOB_TOKEN") == "" {
		errs = append(errs, "neither GITEA_TOKEN nor CI_JOB_TOKEN is set")
	}

	if branch := c.env("SEMREL_PLUGIN_BRANCH"); branch != "" {
		gotBranch := c.env("GITEA_REF_NAME")
		if gotBranch == "" {
			gotBranch = strings.TrimPrefix(c.env("GITEA_REF"), "refs/heads/")
		}
		if gotBranch != branch {
			errs = append(errs, fmt.Sprintf("branch mismatch: want %q got %q", branch, gotBranch))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("%s", strings.Join(errs, "; "))
	}
	return nil
}
