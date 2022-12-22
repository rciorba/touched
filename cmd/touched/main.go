package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// These vars are overriden at compile time by goreleaser
var (
	version = "development"
	commit  = "noop"
)

func PrintCommitted(branch string) error {
	cmd := exec.Command("git", "diff", fmt.Sprintf("%s...HEAD", branch), "--name-only")
	cmd.Env = os.Environ()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

func PrintUncommitted() error {
	cmd := exec.Command("git", "diff", "--name-only")
	cmd.Env = os.Environ()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

func main() {
	var IgnoreUncommitted bool
	var rootCmd = &cobra.Command{
		Use:   "print [string to print]",
		Short: "Print anything to the screen",
		Long: `touched will help you easily list modified files
It assumes you use git for version controll.`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				err := PrintCommitted(args[0])
				if err != nil {
					return err
				}
			}
			if !IgnoreUncommitted {
				err := PrintUncommitted()
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	rootCmd.PersistentFlags().BoolVar(
		&IgnoreUncommitted, "iu", false, "ignore uncommited files",
	)

	rootCmd.Execute()
}
