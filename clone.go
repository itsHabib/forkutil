package forkutil

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CloneCmd represents the command used clone a GitHub repository
var CloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone repository from GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("You must supply a repository to clone")
		}
		if err := CloneRepository(args[0], ref, create); err != nil {
			log.Fatalln("error when cloning repository:", err)
		}
	},
}

// CloneRepository clones a GitHub repository
func CloneRepository(repository, ref string, create bool) error {
	repo, err := NewGHRepo(repository)
	if err != nil {
		return err
	}
	if err := repo.Clone(viper.GetString("location")); err != nil {
		return err
	}
	if err := repo.Checkout(ref, create); err != nil {
		return err
	}
	fmt.Printf("Cloned repository to: %s\n", repo.RepoDir)
	return nil
}

// Reference and Create hold values of flags passed to CloneRepository
var ref string
var create bool

// Initializes the CloneCmd with arguments it can take
func init() {
	CloneCmd.PersistentFlags().StringVar(&ref, "ref", "master",
		"specific reference to check out")
	CloneCmd.PersistentFlags().BoolVar(&create, "create", false,
		"create the reference if it does not exist")
}
