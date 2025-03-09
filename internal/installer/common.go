package installer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/LoftyReign/vine/internal/utils"
)

const startDelimiter = "# Vine stuff"
const endDelimiter = "# ffuts eniV"

func bashrc(overwrite bool) error {

	return nil
}

func runAuth(overwrite bool) error {

	return nil
}

func secretKey(overwrite bool) error {

	return nil
}

func defaultAliasesAndScripts(overwrtie bool) error {

	return nil
}

//////

func getBashrcTemplate() (string, error) {
	var outputString = "\n"

	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return "", fmt.Errorf("Error getting working directory: %e", err)
	}

	file, err := utils.ReadLines(filepath.Join(rootDir, "config", ".bashrc"))
	if err != nil {
		return "", fmt.Errorf("Error reading .bashrc template: %e", err)
	}

	for _, line := range file {
		outputString += line + "\n"
	}

	return outputString, nil
}

func checkForDelimiter(lines []string) bool {
	for _, line := range lines {
		if strings.TrimSpace(line) == strings.TrimSpace(startDelimiter) {
			return true
		}
	}

	return false
}
