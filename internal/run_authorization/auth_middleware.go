package runauthorization

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/LoftyReign/vine/internal/models"
	"github.com/LoftyReign/vine/internal/state"
	"github.com/spf13/cobra"
)

func AuthRun(fn func(cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		allowed, err := authorizeCommand(cmd)
		if err != nil {
			fmt.Printf("Error authorizing command: %s\n", err)
			os.Exit(1)
		}

		if !allowed {
			fmt.Printf("You do not have permission to run this command: %s\n", cmd.CalledAs())
			os.Exit(1)
		}

		fn(cmd, args)
	}
}

func authorizeCommand(cmd *cobra.Command) (bool, error) {
	vinerunAuths := path.Join(state.Location.Auth, "vine_run_auth.json")
	commands := models.AuthorizedCommand{}

	jsonstr, err := os.ReadFile(vinerunAuths)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal([]byte(jsonstr), &commands)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	permission := commands.Command[cmd.CalledAs()]

	valid := checkAuth(permission)

	return valid, nil
}

func checkAuth(permission string) bool {
	return strings.EqualFold(permission, os.Getenv(state.VineVars.VineExecutionVar))
}
