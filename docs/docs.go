package docs

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// DocsCmd represents the command used to get docs from a Github repository
var DocsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Read the documentation for a repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("Must supply repository argument")
		}
		content := GetRepositoryReadme(args[0])
		fmt.Println(content)
	},
}

// GetRepositoryReadme returns a repositories readme
// TODO implement actual function using Github API
func GetRepositoryReadme(repository string) string {
	return repository
}
