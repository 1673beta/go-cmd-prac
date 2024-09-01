package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var defaultRepo = "origin"
var defaultBranch = "master"
var lowMemory bool
var branch string
var repo string

var updateCmd = &cobra.Command{
	Use:   "update [repo] [branch]",
	Short: "Update misskey to latest version for systemd.",
	Run: func(cmd *cobra.Command, args []string) {
		out, err := exec.Command("git", "pull", repo, branch).CombinedOutput()
		if err != nil {
			fmt.Printf("error while git pull: %v, output: %s\n", err, out)
			os.Exit(1)
		}
		fmt.Println(string(out))

		out, err = exec.Command("git", "submodule", "update", "--init").CombinedOutput()
		if err != nil {
			fmt.Printf("error while submodule update: %v, output: %s", err, out)
			os.Exit(1)
		}
		fmt.Println(string(out))

		if lowMemory {
			os.Setenv("NODE_OPTIONS", "--max_old_space_size=3072")
		}

		os.Setenv("NODE_ENV", "production")

		out, err = exec.Command("pnpm", "install", "--frozen-lockfile", "--non-interactive").Output()
		if err != nil {
			fmt.Println("error while installing module:", err)
			os.Exit(1)
		}
		fmt.Println(string(out))

		out, err = exec.Command("pnpm", "run", "build").Output()
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		fmt.Println(string(out))

		out, err = exec.Command("pnpm", "run", "migrate").Output()
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		fmt.Println(string(out))
	},
}

func init() {
	updateCmd.Flags().BoolVarP(&lowMemory, "low-memory", "l", false, "Enable low memory mode.")
	updateCmd.Flags().StringVarP(&branch, "branch", "b", defaultBranch, "Speciy branch to pull.")
	updateCmd.Flags().StringVarP(&repo, "repo", "r", defaultRepo, "Specify repository to pull.")
	rootCmd.AddCommand(updateCmd)
}
