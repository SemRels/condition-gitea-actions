package plugin

import (
	"strings"
	"testing"
)

func env(kv map[string]string) func(string) string {
	return func(key string) string { return kv[key] }
}

func TestCheck_HappyPath(t *testing.T) {
	t.Parallel()

	err := NewWithEnv(env(map[string]string{
		"GITEA_ACTIONS": "true",
		"GITEA_TOKEN":   "token",
	})).Check()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCheck_AllowsCIJobToken(t *testing.T) {
	t.Parallel()

	err := NewWithEnv(env(map[string]string{
		"GITEA_ACTIONS": "true",
		"CI_JOB_TOKEN":  "token",
	})).Check()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCheck_RequiresActions(t *testing.T) {
	t.Parallel()

	err := NewWithEnv(env(map[string]string{"GITEA_TOKEN": "token"})).Check()
	if err == nil || !strings.Contains(err.Error(), "GITEA_ACTIONS") {
		t.Fatalf("expected actions error, got: %v", err)
	}
}

func TestCheck_RequiresToken(t *testing.T) {
	t.Parallel()

	err := NewWithEnv(env(map[string]string{"GITEA_ACTIONS": "true"})).Check()
	if err == nil || !strings.Contains(err.Error(), "GITEA_TOKEN") {
		t.Fatalf("expected token error, got: %v", err)
	}
}

func TestCheck_BranchFromRefName(t *testing.T) {
	t.Parallel()

	err := NewWithEnv(env(map[string]string{
		"GITEA_ACTIONS":        "true",
		"GITEA_TOKEN":          "token",
		"SEMREL_PLUGIN_BRANCH": "main",
		"GITEA_REF_NAME":       "main",
	})).Check()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCheck_BranchFromRefFallback(t *testing.T) {
	t.Parallel()

	err := NewWithEnv(env(map[string]string{
		"GITEA_ACTIONS":        "true",
		"GITEA_TOKEN":          "token",
		"SEMREL_PLUGIN_BRANCH": "main",
		"GITEA_REF":            "refs/heads/main",
	})).Check()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCheck_BranchMismatch(t *testing.T) {
	t.Parallel()

	err := NewWithEnv(env(map[string]string{
		"GITEA_ACTIONS":        "true",
		"GITEA_TOKEN":          "token",
		"SEMREL_PLUGIN_BRANCH": "main",
		"GITEA_REF_NAME":       "develop",
	})).Check()
	if err == nil || !strings.Contains(err.Error(), "branch mismatch") {
		t.Fatalf("expected branch mismatch, got: %v", err)
	}
}

func TestCheck_MultipleErrors(t *testing.T) {
	t.Parallel()

	err := NewWithEnv(env(map[string]string{})).Check()
	if err == nil || !strings.Contains(err.Error(), "GITEA_ACTIONS") || !strings.Contains(err.Error(), "GITEA_TOKEN") {
		t.Fatalf("expected combined errors, got: %v", err)
	}
}
