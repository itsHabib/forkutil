package forkutil

import "testing"

// TestSearchByKeyword ensures searching GitHub by keyword functions properly
// TODO: Implement Real tests for GitHub API
func TestSearchByKeyword(t *testing.T) {
	repositoryList := SearchByKeyword([]string{"first", "second"})
	if repositoryList[0] != "exampleRepo" {
		t.Errorf("Error searching by keyword, expected=%s, got=%s",
			"exampleRepo", repositoryList[0])
	}
}
