package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/charmbracelet/glamour"
	flag "github.com/spf13/pflag"
)

var (
	// AppName is the name of the application
	AppName = "mdcat"
	// Author is the author of the application
	Author = "@mrinjamul"
	// Version is the version of the application
	Version = "0.1.0"
	// CommitHash is the commit hash of the application
	CommitHash = "none"
	// BuildDate is the date of the build
	BuildDate = "unknown"
)

// flag variables
var (
	flagHelp    bool
	flagVersion bool
)

func main() {
	// parse flags
	flag.Parse()

	// if user does not supply flags, print usage
	//	if flag.NFlag() == 0 {
	//		printUsage()
	//	}

	if flagHelp {
		printUsage()
		os.Exit(0)
	}

	if flagVersion {
		printVersion()
		os.Exit(0)
	}

	// implement cat command
	args := flag.Args()

	// check if argument is present
	if len(args) == 0 {
		fmt.Println("No file specified")
		os.Exit(1)
	}
	// check if argument is greater than 1
	if len(args) > 1 {
		fmt.Println("Only one file can be specified")
		os.Exit(1)
	}
	// check if file exists
	if _, err := os.Stat(args[0]); os.IsNotExist(err) {
		fmt.Println("File does not exist")
		os.Exit(1)
	}
	// check if it is a file else print error
	if isDir(args[0]) {
		fmt.Println("File is a directory")
		os.Exit(1)
	}
	if !isMarkdownFile(args[0]) {
		printFiles(args[0])
		os.Exit(0)
	}

	printMarkdownFile(args[0])

}

func init() {
	flag.BoolVarP(&flagHelp, "help", "h", false, "print help")
	flag.BoolVarP(&flagVersion, "version", "v", false, "print version")
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", AppName)
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(0)
}
func printVersion() {
	fmt.Println(AppName + " " + Version + "+" + CommitHash + " " + BuildDate + " by " + Author)
}

// isDir checks if a path is a directory
func isDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// check if the file is markdown file or not using extension
func isMarkdownFile(path string) bool {
	ext := filepath.Ext(path)
	if ext == ".md" || ext == ".markdown" || ext == ".mdown" || ext == ".mkd" || ext == ".mkdn" || ext == ".mdwn" || ext == ".mdtxt" || ext == ".mdtext" {
		return true
	}
	return false
}

// printFiles get the file path and print the content into stdout
func printFiles(path string) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(buf))
}

// printMarkdownFile get the file path and print the markdown with the help of glamour package
func printMarkdownFile(path string) {
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
