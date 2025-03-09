package cmd

import (
	runauthorization "github.com/LoftyReign/vine/internal/run_authorization"
	parser "github.com/LoftyReign/vine/internal/yml_parser"
	"github.com/spf13/cobra"
)

var alias string

var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "runs the thing",
	Args:  cobra.ExactArgs(1),
	Run: runauthorization.AuthRun(
		func(cmd *cobra.Command, args []string) {
			parser.Alias(args[0])
		}),
}

func init() {
	rootCmd.AddCommand(aliasCmd)
}
