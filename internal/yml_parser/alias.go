package parser

import (
	"fmt"

	"github.com/LoftyReign/vine/internal/utils"
)

func Alias(alias string) {
	vault, err := getVault()
	if err != nil {
		fmt.Printf("Error: %e", err)
		return
	}

	for _, c := range vault.Commands {
		if utils.Contains(c.Aliases, alias) {
			fmt.Printf(c.Command)
			return
		}
	}
}
