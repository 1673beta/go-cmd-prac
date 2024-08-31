package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var nyanCount int

var nyanCmd = &cobra.Command{
	Use:   "nyan",
	Short: "nyaaan!",
	Run: func(cmd *cobra.Command, args []string) {
		nyanCount++
		if nyanCount == 15 {
			fmt.Println("There are no Easter Eggs in this program.")
		} else if nyanCount == 17 {
			fmt.Println("There really are no Easter Eggs in this program.")
		} else if nyanCount == 18 {
			fmt.Println("Didn't I tell you that there are no Easter Eggs in this program?")
		} else if nyanCount == 19 {
			fmt.Println("Stop it! ")
		} else if nyanCount == 20 {
			fmt.Println("Okay, okaym if I give you an Easter Egg, will you go away?")
		} else if nyanCount == 21 {
			fmt.Println("All right, you win.")
		} else if nyanCount == 22 {
			fmt.Println("What is it? It's an cute seeking cat!")
		} else if nyanCount >= 23 {
			fmt.Println("...There is no cats. They took off when you took your eye off them.")
			nyanCount = 0
		} else {
			fmt.Println("nyan!")
		}
	},
}

func init() {
	rootCmd.AddCommand(nyanCmd)
}
