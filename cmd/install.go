package cmd

import (
	"github.com/LoftyReign/vine/internal/installer"
	runauthorization "github.com/LoftyReign/vine/internal/run_authorization"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "installs vine",
	Long:  "creates the necessary directories and adds the vine function to the .bashrc",
	Run: runauthorization.AuthRun(
		func(cmd *cobra.Command, args []string) {
			installer.Install()
		}),
}

func init() {
	rootCmd.AddCommand(installCmd)
}
