package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Review the current code changes",
	Run: func(cmd *cobra.Command, args []string) {
		gitCmd := exec.Command(
			"git",
			"--no-pager",
			"diff",
			"--name-only",
			"--cached",
			"--merge-base",
			"origin/main",
		)
		stdout, err := gitCmd.Output()
		if err != nil {
			log.Fatal(err)
		}

		changedFiles := string(stdout)

		for file := range strings.SplitSeq(changedFiles, "\n") {
			fmt.Print(file)
		}
	},
}

func init() {
	rootCmd.AddCommand(reviewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reviewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reviewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
