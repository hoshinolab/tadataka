package util

import (
	"bufio"
	"fmt"
	"os"
)

func CLIQuestion() bool {
	result := true
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		if input == "Y" || input == "y" {
			break
		} else if input == "N" || input == "n" {
			result = false
			break
		} else {
			fmt.Println("Please input y or n. Try again.")
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return result
}
