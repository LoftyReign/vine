package main

import (
	"fmt"

	"github.com/LoftyReign/vine/cmd"
	"github.com/LoftyReign/vine/internal/models"
	"github.com/LoftyReign/vine/internal/state"
	"github.com/LoftyReign/vine/internal/utils"
)

func main() {
	location := &models.Location{
		Path:    "~/.vine/",
		BashRC:  "~/.bashrc",
		Scripts: "~/.vine/vine_scripts/",
		Auth:    "~/.vine/.auth/",
		Config:  "~/.vine/vine_aliases.yaml",
	}
	state.SetLocation(location)

	vineVars, err := utils.GetVineVars()
	if err != nil {
		fmt.Printf("Error setting vine vars: %s", err)
	}
	state.SetVineVars(vineVars)

	cmd.Execute()
}
