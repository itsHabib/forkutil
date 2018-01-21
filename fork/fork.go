package fork

import (
	"log"

	"github.com/spf13/cobra"
)

// ForkCmd is the command responsible for forking GitHub repositories
var ForkCmd = &cobra.Command{
	Use:   "fork",
	Short: "fork a GitHub repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("You must provide a repository")
		}
		if err := ForkRepository(args[0]); err != nil {
			log.Fatalln("Unable to fork repository:", err)
		}
	},
}

// ForkRepository forks a GitHub repository
func ForkRepository(repository string) error {
	return nil
}
