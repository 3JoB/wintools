// A simple grep program written in golang
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/grafana/regexp"
)

// Define the command-line flags
var (
	ignoreCase = flag.Bool("i", false, "ignore case")
	invert     = flag.Bool("v", false, "invert match")
	recursive  = flag.Bool("r", false, "recursive search")
	help       = flag.Bool("h", false, "show help")
)

// Define the usage message
var usage = `Usage: grep [options] pattern [files...]
Options:
  -i    ignore case
  -v    invert match
  -r    recursive search
  -h    show help
`

// Define a global variable to store the pattern
var pattern string

// Define a function to print the usage message and exit
func printUsage() {
	fmt.Fprint(os.Stderr, usage)
	os.Exit(1)
}

// Define a function to print an error message and exit
func printError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(2)
}

// Define a function to search a file for the pattern and print the matching lines
func grepFile(file *os.File) {
	// Create a buffered reader for the file
	reader := bufio.NewReader(file)

	// Create a regular expression from the pattern
	var re *regexp.Regexp
	var err error
	if *ignoreCase {
		re, err = regexp.Compile("(?i)" + pattern)
	} else {
		re, err = regexp.Compile(pattern)
	}
	if err != nil {
		printError(err)
	}

	// Loop through the lines of the file
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				printError(err)
			}
			break // end of file
		}

		// Check if the line matches the pattern
		match := re.MatchString(line)

		// If invert flag is set, reverse the match result
		if *invert {
			match = !match
		}

		// If match is true, print the line with the file name
		if match {
			fmt.Print(line)
		}
	}
}

// Define a function to search a directory for the pattern recursively
func grepDir(dir string) {
	// Open the directory and read its contents
	d, err := os.Open(dir)
	if err != nil {
		printError(err)
	}
	defer d.Close()
	files, err := d.Readdir(-1)
	if err != nil {
		printError(err)
	}

	// Loop through the files and directories in the directory
	for _, file := range files {
		path := dir + string(os.PathSeparator) + file.Name()

		// If it is a file, grep it
		if file.Mode().IsRegular() {
			f, err := os.Open(path)
			if err != nil {
				printError(err)
			}
			defer f.Close()
			grepFile(f)
		}

		// If it is a directory and recursive flag is set, grep it recursively
		if file.Mode().IsDir() && *recursive {
			grepDir(path)
		}
	}
}

// Define the main function
func main() {
	// Parse the command-line flags
	flag.Parse()

	// If help flag is set or no arguments are given, print the usage message and exit
	if *help || flag.NArg() == 0 {
		printUsage()
	}

	// Get the pattern from the first argument
	pattern = flag.Arg(0)

	// If no files are given, grep the standard input
	if flag.NArg() == 1 {
		grepFile(os.Stdin)
		return
	}

	// Loop through the files given as arguments
	for _, name := range flag.Args()[1:] {
		// If the name is "-", grep the standard input
		if name == "-" {
			grepFile(os.Stdin)
			continue
		}

		// Get the file info and check if it is a directory or a file
		info, err := os.Stat(name)
		if err != nil {
			printError(err)
		}
		if info.IsDir() {
			// If it is a directory, grep it recursively
			grepDir(name)
		} else {
			// If it is a file, grep it
			file, err := os.Open(name)
			if err != nil {
				printError(err)
			}
			defer file.Close()
			grepFile(file)
		}
	}
}
