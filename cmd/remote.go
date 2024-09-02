package cmd

import (
	"fmt"
	"mkctl/cmd/util"
	"os"

	"github.com/spf13/cobra"
)

var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Manage remote servers settings",
}

var suspendCmd = &cobra.Command{
	Use:   "suspend",
	Short: "Suspend remote server.",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := util.ConnectToDb()
		if err != nil {
			fmt.Printf("error while connecting to database: %v", err)
			os.Exit(1)
		}

		host := args[0]

		query := `UPDATE instance SET suspensionState = 'manuallySuspended' WHERE instance."host" = $1`
		_, err = db.Exec(query, host)
		if err != nil {
			fmt.Printf("error while suspending remote server: %v", err)
			os.Exit(1)
		}
	},
}

var unsuspendCmd = &cobra.Command{
	Use:   "unsuspend",
	Short: "Unsuspend remote server.",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := util.ConnectToDb()
		if err != nil {
			fmt.Printf("error while connecting to database: %v", err)
			os.Exit(1)
		}

		host := args[0]

		query := `UPDATE instance SET suspensionState = 'none' WHERE instance."host" = $1`
		_, err = db.Exec(query, host)
		if err != nil {
			fmt.Printf("error while unsuspending remote server: %v", err)
			os.Exit(1)
		}
	},
}

var goneCmd = &cobra.Command{
	Use:   "gone",
	Short: "Set remote server as deleted.",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := util.ConnectToDb()
		if err != nil {
			fmt.Printf("error while connecting to database: %v", err)
			os.Exit(1)
		}

		host := args[0]

		query := `UPDATE instance SET suspensionState = 'goneSuspended' WHERE instance."host" = $1`
		_, err = db.Exec(query, host)
		if err != nil {
			fmt.Printf("error while setting remote server as gone: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(remoteCmd)
	remoteCmd.AddCommand(suspendCmd)
	remoteCmd.AddCommand(unsuspendCmd)
	remoteCmd.AddCommand(goneCmd)
}
