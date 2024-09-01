package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var nyanCmd = &cobra.Command{
	Use:   "nyan [flags]",
	Short: "nyaaan!",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			switch args[0] {
			case "nyan":
				fmt.Println("nyan!")
			case "nyannyan":
				fmt.Println("nyannyan!")
			case "22":
				fmt.Println("There really are no Easter eggs in this progoram.")
			case "222":
				fmt.Println("Didn't I tell you there are no Easter eggs in this program?")
			case "2222":
				fmt.Println("Stop it!")
			case "22222":
				fmt.Println("Okay, okay, if I give you an Easter egg, will you go away?")
			case "222222":
				fmt.Println(`
				All right, you win.

                               /^--^\
                       -------/      \
                      /               \
                     /                |
     /--------------/                  --------\
   /----------------------------------------------
				`)
			case "2222222":
				fmt.Println("What is it? It's a cat on the carpet, of course.")
			case "22222222":
				fmt.Println("Cat ran away when you took your eyes off it.")
			default:
				fmt.Println("nyan?")
			}
		} else {
			fmt.Println("nya~")
		}
	},
}

func init() {
	rootCmd.AddCommand(nyanCmd)
}
