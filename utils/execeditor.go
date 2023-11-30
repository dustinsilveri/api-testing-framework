package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecEditor() {
	cmd := exec.Command("bg", "") // need to finish.

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
