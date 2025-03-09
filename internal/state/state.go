package state

import (
	"log"
	"os"
	"strings"

	"github.com/LoftyReign/vine/internal/models"
)

var Location *models.Location
var HomeDir string
var VineVars *models.VineVars

func SetLocation(location *models.Location) {
	if HomeDir == "" {
		setHomeDir()
	}
	Location = &models.Location{
		Path:    strings.ReplaceAll(location.Path, "~", HomeDir),
		BashRC:  strings.ReplaceAll(location.BashRC, "~", HomeDir),
		Config:  strings.ReplaceAll(location.Config, "~", HomeDir),
		Auth:    strings.ReplaceAll(location.Auth, "~", HomeDir),
		Scripts: strings.ReplaceAll(location.Scripts, "~", HomeDir),
	}
}

func SetVineVars(vineVars *models.VineVars) {
	VineVars = vineVars
}

func setHomeDir() {
	var err error
	HomeDir, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
}
