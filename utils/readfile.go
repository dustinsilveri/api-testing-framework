package utils

import (
	"fmt"
	"os"
)

/*
Provides functionality to read and parse local files.
*/

func ReadFile(filename string) ([]byte, error) {
	b, err := os.ReadFile(filename) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	return b, err
}
