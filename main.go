package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/mrinjamul/mdcat/app"
	flag "github.com/spf13/pflag"
)

var (
	// AppName is the name of the application
	AppName = "mdcat"
	// Author is the author of the application
	Author = "@mrinjamul"
	// Version is the version of the application
	Version = "dev"
	// CommitHash is the commit hash of the application
	CommitHash = "none"
	// BuildDate is the date of the build
	BuildDate = "unknown"
)

// flag variables
var (
	flagHelp    bool
	flagVersion bool
	flagRaw     bool
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
	if len(args) == 0 && !app.IsInputFromPipe() {
		fmt.Println("No file specified")
		os.Exit(1)
	}
	// check if data is piped
	if app.IsInputFromPipe() {
		// read from pipe
		reader := bufio.NewReader(os.Stdin)
		var output []rune

		for {
			input, _, err := reader.ReadRune()
			if err != nil && err == io.EOF {
				break
			}
			output = append(output, input)
		}
		// print from pipe
		if flagRaw {
			fmt.Println(string(output))
			return
		}
		out, err := glamour.Render(string(output), "dark")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(out)
		return
	}
	for _, path := range args {
		// check if file exists
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("File does not exist")
			continue
		}
		// check if it is a file else print error
		if app.IsDir(path) {
			fmt.Println("File is a directory")
			continue
		}
		// check if it is markdown file
		if !app.IsMarkdownFile(path) {
			app.PrintFiles(path)
			continue
		}
		if flagRaw {
			app.PrintFiles(path)
			continue
		}
		app.PrintMarkdownFile(path)
	}

}

func init() {
	flag.BoolVarP(&flagRaw, "raw", "r", false, "print raw markdown")
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
