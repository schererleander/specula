//go:build unix
// +build unix
package main

import (
	"fmt"
	"os"
	//tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	filename string
	size int64
	created string
	lastModified string
	owner string
	permission string
	description string
	width int
	height int
	path string
	error string
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Error, invalid arguments: spectula [path]")
		return
	}

	path := args[0]
	info, err := os.Stat(path)
	if err != nil {
		fmt.Printf("Error accessing %s: %v\n", path, err)
		os.Exit(1)
	}

	model := &model{path: path}
	if !info.Mode().IsRegular() {
		fmt.Printf("Error, %s is not a regular file\n", path)
		os.Exit(1)
	}

	model.populateFromInfo(info)
	model.getDescription(path)

	fmt.Println(model)
}
