package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "web-cli",
	Short: "A CLI news reader for atom feeds",
	Long: `Read atom feeds (e.g. heise online newsfeed)
and show details on the command line.`,
}

var cmdLs = &cobra.Command{
	Use:   "ls",
	Short: "List articles",
	Long:  `Show the first 5 articles from heise.de`,
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

var cmdDescribe = &cobra.Command{
	Use:   "describe [id]",
	Short: "Show details for an article",
	Long:  `Show details for an article`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		describe(args[0])
	},
}

func Exec() {
	rootCmd.AddCommand(cmdLs)
	rootCmd.AddCommand(cmdDescribe)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
