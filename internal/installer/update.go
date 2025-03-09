package installer

import (
	"fmt"
)

func Update() {
	fmt.Println("You are about to update the local references to the repo")
	getVinePart()
}

func getVinePart() error {
	// var outputString = ""
	//
	// file, err := utils.ReadLines(state.Location.BashRC)
	// if err != nil {
	// 	return fmt.Errorf("Error reading .bashrc: %e", err)
	// }
	// fmt.Println(file)
	//
	// startDelimiterFound := false
	// for _, line := range file {
	// 	if strings.EqualFold(line, startDelimiter) {
	// 		startDelimiterFound = true
	// 	}
	// 	if startDelimiterFound {
	// 		outputString += line
	// 	}
	// 	if strings.EqualFold(line, endDelimiter) {
	// 		break
	// 	}
	// }

	return nil
}
