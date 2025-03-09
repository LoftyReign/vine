package installer

import (
	"fmt"
	"os"

	"github.com/LoftyReign/vine/internal/state"
)

func Install() {
	// lines, err := utils.ReadLines(state.Location.BashRC)
	// if err != nil {
	// 	fmt.Printf("Error reading the lines: %v", err)
	// 	return
	// }
	//
	// if checkForDelimiter(lines) {
	// 	fmt.Printf(".bashrc already contains vine\n")
	// 	return
	// }
	//
	// err = addVine()
	// if err != nil {
	// 	fmt.Printf("Error adding vine to .bashrc: %s", err)
	// 	return
	// }
	// fmt.Println("Vine added to .bashrc")

	fmt.Println("Install successfully called")
}

func addVine() error {
	file, err := os.OpenFile(state.Location.BashRC, os.O_APPEND|os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	bashrcTemplate, err := getBashrcTemplate()
	if err != nil {
		return err
	}

	_, err = file.WriteString(bashrcTemplate)
	if err != nil {
		return fmt.Errorf("error writing to file: %e", err)
	}

	return nil
}
