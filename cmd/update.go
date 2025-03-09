package cmd

import (
	"github.com/LoftyReign/vine/internal/installer"
	runauthorization "github.com/LoftyReign/vine/internal/run_authorization"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update installed dependencies",
	Long:  "updated the bashrc and vine_aly",
	Run: runauthorization.AuthRun(
		func(cmd *cobra.Command, args []string) {
			installer.Update()
		}),
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
