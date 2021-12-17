package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/charmbracelet/glamour"
)

// IsInputFromPipe returns true if the input is from a pipe
func IsInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

// IsDir checks if a path is a directory
func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// IsMarkdownFile check if the file is markdown file or not using extension
func IsMarkdownFile(path string) bool {
	ext := filepath.Ext(path)
	if ext == ".md" || ext == ".markdown" || ext == ".mdown" || ext == ".mkd" || ext == ".mkdn" || ext == ".mdwn" || ext == ".mdtxt" || ext == ".mdtext" {
		return true
	}
	return false
}

// PrintFiles get the file path and print the content into stdout
func PrintFiles(path string) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(buf))
}

// PrintMarkdownFile get the file path and print the markdown with the help of glamour package
func PrintMarkdownFile(path string) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	out, err := glamour.Render(string(buf), "dark")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(out)
}
