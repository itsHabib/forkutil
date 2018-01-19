package search

import (
	"fmt"

	"github.com/spf13/cobra"
)

// SearchCmd represents the command used to search GitHub repositories
var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for GitHub repositories by keyword",
	Run: func(cmd *cobra.Command, args []string) {
		repositoryList := SearchByKeyword(args)
		for _, repository := range repositoryList {
			fmt.Println(repository)
		}
	},
}

// SearchByKeyword searches GitHub repositories by keyword
func SearchByKeyword(keywords []string) []string {
	return []string{"exampleRepo"}
}
