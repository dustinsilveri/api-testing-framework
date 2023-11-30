package utils

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func ShowModules(Modules []string) {
	// When show modules is called, print the list of all imported modules
	fmt.Println("")
	fmt.Println("Modules")
	filler := strings.Repeat("=", utf8.RuneCountInString("Modules")) // fill - under the heading
	fmt.Println(filler)
	fmt.Println("")

	for i, item := range Modules {
		fmt.Println(i, item)

	}

	fmt.Println("")
}

func ShowTemplates(Templates []string) {
	// When show templates is called, print the list of all imported templates
	fmt.Println("")
	fmt.Println("Templates")
	filler := strings.Repeat("=", utf8.RuneCountInString("Templates")) // fill - under the heading
	fmt.Println(filler)
	fmt.Println("")

	for i, item := range Templates {
		fmt.Println(i, item)
	}

	fmt.Println("")

}
