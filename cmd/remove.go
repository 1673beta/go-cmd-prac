package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"mkctl/cmd/util"
)

var deleteDays int
var remoteOnly bool

type DbConfig struct {
	DB struct {
		Host  string `yaml:"host"`
		Port  int    `yaml:"port"`
		Db    string `yaml:"db"`
		User  string `yaml:"user"`
		Pass  string `yaml:"pass"`
		Extra struct {
			SSL bool `yaml:"ssl"`
		} `yaml:"extra,omitempty"`
	} `yaml:"db"`
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove old posts from database.",
	Run: func(cmd *cobra.Command, args []string) {
		cutoff := time.Now().AddDate(0, 0, -deleteDays)
		ms := cutoff.Sub(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)).Milliseconds()

		encoded := strconv.FormatInt(ms, 36)

		db, err := util.ConnectToDb()
		if err != nil {
			fmt.Printf("error while connecting to database: %v", err)
			os.Exit(1)
		}

		if remoteOnly {
			query := fmt.Sprintf(`DELETE FROM note WHERE SUBSTRING(note."id", 1, 8) <= '%s' AND note."userHost" IS NOT NULL`, encoded)
			_, err := db.Exec(query)
			if err != nil {
				fmt.Printf("error while deleting old notes: %v", err)
				os.Exit(1)
			}
		} else {
			query := fmt.Sprintf(`DELETE FROM note WHERE SUBSTRING(note."id", 1, 8) <= '%s`, encoded)
			_, err := db.Exec(query)
			if err != nil {
				fmt.Printf("error while deleting old notes: %v", err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	updateCmd.Flags().IntVarP(&deleteDays, "days", "d", 120, "How old notes have to be before they are deleted. Defaults to 120.")
	updateCmd.Flags().BoolVarP(&remoteOnly, "remote", "r", false, "Only delete notes from remote server.")
	rootCmd.AddCommand(removeCmd)
}
