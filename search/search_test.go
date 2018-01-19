package search

import "testing"

// TestSearchByKeyword ensures searching Github by keyword functions properly
// TODO: Implement Real tests for Github API
func TestSearchByKeyword(t *testing.T) {
	repositoryList := SearchByKeyword([]string{"first", "second"})
	if repositoryList[0] != "exampleRepo" {
		t.Errorf("Error searching by keyword, expected=%s, got=%s",
			"exampleRepo", repositoryList[0])
	}
}
