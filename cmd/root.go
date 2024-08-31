package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mkctl",
	Short: "USAGE: mkctl [command]",
	Long:  `mkctl is a CLI tool for managing Misskey server. This command is a root command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from mkctl!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
