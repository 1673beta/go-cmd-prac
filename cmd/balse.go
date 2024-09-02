package cmd

import (
	"bufio"
	"fmt"
	"mkctl/cmd/util"
	"os"
	"strings"
	"time"

	"github.com/cheggaaa/pb"
	"github.com/spf13/cobra"
)

var balseCmd = &cobra.Command{
	Use:   "balse",
	Short: "If you run this command, your database record will be delete.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("I'll give you 3 minutes starting now.")
		fmt.Println("\033[31mWarning: This command will delete all records in the database.\033[0m")

		bar := pb.StartNew(180)

		for i := 0; i < 180; i++ {
			bar.Increment()
			time.Sleep(time.Second)
		}

		bar.Finish()

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Time's up. What do you say? (balse/no)")
		yes, _ := reader.ReadString('\n')

		if strings.TrimSpace(yes) == "balse" {
			fmt.Println("My eyes, my eyes...")
			db, err := util.ConnectToDb()
			if err != nil {
				fmt.Printf("error while connecting to database: %v", err)
				os.Exit(1)
			}

			dbConfig, err := util.LoadConfig()
			if err != nil {
				fmt.Printf("error while loading config: %v", err)
				os.Exit(1)
			}

			query := `DROP DATABASE $1`
			_, err = db.Exec(query, dbConfig.DB.Db)
			if err != nil {
				fmt.Printf("error while deleting database: %v", err)
			}
		} else if strings.TrimSpace(yes) == "no" {
			fmt.Println("Just where will you go?")
		}
	},
}

func init() {
	rootCmd.AddCommand(balseCmd)
}
