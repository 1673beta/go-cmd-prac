package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of mkctl",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mkctl version: ", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
