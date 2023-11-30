package utils

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

func InitArrays(Modules []string, dir string) []string {
	var Paths []string
	Paths, err := FilePathWalkDir(dir)
	//sort.Strings(Paths)
	if err != nil {
		fmt.Println(err)
	}

	if dir != "templates" {
		for _, file := range Paths {
			fn := FilenameWithoutExtension(file)    // strip file extensions, except if templates dir, then show them.
			fn = strings.Replace(fn, "\\", "/", -1) // fix windows directory \ to standarize
			Modules = append(Modules, fn)           // Add each path in the modules dir to array.
		}
	} else {
		Modules = append(Modules, Paths...) // Add each path in the modules dir to array.
	}

	return Modules
}

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		sort.Strings(files)
		return nil
	})
	return files, err
}

func FilenameWithoutExtension(name string) string {
	return strings.TrimSuffix(name, path.Ext(name))
}
