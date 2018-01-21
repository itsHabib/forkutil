package main

import (
	"os"

	"github.com/itsHabib/forkutil/clone"
	"github.com/itsHabib/forkutil/docs"
	"github.com/itsHabib/forkutil/fork"
	"github.com/itsHabib/forkutil/search"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/// Root Command used to work with other sub commands
var rootCmd = &cobra.Command{}

func main() {
	rootCmd.Execute()
}

func init() {
	rootCmd = &cobra.Command{
		Use:   "forkutil",
		Short: "Project Forking tool For GitHub",
	}
	rootCmd.AddCommand(search.SearchCmd)
	rootCmd.AddCommand(docs.DocsCmd)
	rootCmd.AddCommand(clone.CloneCmd)
	rootCmd.AddCommand(fork.ForkCmd)
	viper.SetDefault("location", os.Getenv("HOME"))
	viper.SetConfigName("forkutil")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
}
