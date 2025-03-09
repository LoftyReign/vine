package utils

import (
	"encoding/json"
	"os"
	"path"

	"github.com/LoftyReign/vine/internal/models"
	"github.com/LoftyReign/vine/internal/state"
)

func GetVineVars() (*models.VineVars, error) {
	vine_exposed_config := path.Join(state.Location.Path, ".config", "vine_exposed_config.json")
	var vine_vars models.VineVars

	jsonstr, err := os.ReadFile(vine_exposed_config)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonstr), &vine_vars)
	if err != nil {
		return nil, err
	}

	return &vine_vars, nil
}
