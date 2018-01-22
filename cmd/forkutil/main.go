package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/itsHabib/forkutil"
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
	rootCmd.AddCommand(forkutil.SearchCmd)
	rootCmd.AddCommand(forkutil.DocsCmd)
	rootCmd.AddCommand(forkutil.CloneCmd)
	rootCmd.AddCommand(forkutil.ForkCmd)
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No config file found")
	}
	viper.SetDefault("location", os.Getenv("HOME"))
}
