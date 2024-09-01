package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var defaultRepo = "https://github.com/misskey-dev/misskey.git"
var defaultBranch = "master"
var lowMemory bool

var updateCmd = &cobra.Command{
	Use:   "update [repo] [branch]",
	Short: "Update misskey to latest version for systemd.",
	Run: func(cmd *cobra.Command, args []string) {
		repo := defaultRepo
		branch := defaultBranch
		if len(args) > 0 {
			repo = args[0]
		}
		if len(args) > 1 {
			branch = args[1]
		}
		out, err := exec.Command("git", "pull", repo, branch).Output()
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		fmt.Println(string(out))

		out, err = exec.Command("git", "submodule", "update", "--init").Output()
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		fmt.Println(string(out))

		if lowMemory {
			os.Setenv("NODE_OPTIONS", "--max_old_space_size=3072")
		}

		os.Setenv("NODE_ENV", "production")

		out, err = exec.Command("pnpm", "install", "--frozen-lockfile").Output()
		if err != nil {
			fmt.Println("error:", err)
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
	rootCmd.AddCommand(updateCmd)
}
