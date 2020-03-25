package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "web-cli",
	Short: "Just a demo",
	Long: `Just a small CLI application
	Nice
	One`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TEST")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
