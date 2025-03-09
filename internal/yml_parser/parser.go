package parser

import (
	"fmt"

	"github.com/LoftyReign/vine/internal/models"
	"gopkg.in/yaml.v3"
)

func getVault() (*models.Vault, error) {
	var vault models.Vault

	yamlContent, err := loadAliases()
	if err != nil {
		return nil, fmt.Errorf("Error: %e", err)
	}

	err = yaml.Unmarshal([]byte(yamlContent), &vault)
	if err != nil {
		return nil, fmt.Errorf("Error: %e", err)
	}

	return &vault, nil
}
