package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "get-time",
	Short: "get-time manages and summarizes your timesheets",
}

func Execute() {
	cmd, _, err := rootCmd.Find(os.Args[1:])

	// use default `times` command if none given
	wasGivenCommand := err != nil || cmd.Use != rootCmd.Use || cmd.Flags().Parse(os.Args[1:]) == pflag.ErrHelp

	if !wasGivenCommand {
		args := append([]string{"times"}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
