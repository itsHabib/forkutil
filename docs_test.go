package forkutil

import (
	"testing"
)

func TestGetRepositoryReadme(t *testing.T) {
	content := GetRepositoryReadme("repo")
	if content != "repo" {
		t.Errorf("Error getting repository readme, expected=%s, got=%s",
			"repo", content)
	}
}
