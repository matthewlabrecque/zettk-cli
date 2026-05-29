/*
Copyright © 2026 Matthew Labrecque <mlabrecque2002@gmail.com>

*/
package cmd

import (
	"fmt"
	"strings"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [args]",
	Short: "Search through the zettlekasten",
	Long: `Search through the zettlekasten for a file provided by the argument.
Returns the file name, zettlekasten ID, date created, and date
last modified.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Failed to find user home directory")
		}
		targetDir, err := cmd.Flags().GetString("target")
		if err != nil { fmt.Println(err) }

		// Determine if a target directory has been specified
		// This also performs error checking for valid argument
		var searchDir string
		switch strings.ToUpper(targetDir){
			case "INBOX":
				searchDir = "00-INBOX"
			case "ARCHIVE":
				searchDir = "01-ARCHIVE"
			case "REFERENCES":
				searchDir = "02-REFERENCES"
			default:
				fmt.Println("Error: Invalid argument")
				return
		}	

		// Build the path to search within
		targetPath := filepath.Join(homeDir, "zettlekasten", searchDir)

		// Perform a fuzzy search using the argument passed
		// Strange edge-case: "and" and "the" aren't picked up
		foundFiles := fileTraversal(targetPath, filepath.Clean(args[0]))

		// See how many files are in the list,
		// If there's more than one prompt for the user
		var targFile string = ""
		if len(foundFiles) == 0 {
			fmt.Println("No files found matching search query.")
			return
		} else if len(foundFiles) == 1 {
			targFile = foundFiles[0]
		} else {
			targFile = findFile(foundFiles, filepath.Clean(args[0]))	
		}

		// Build the file path to grab info from
		file, _ := os.Stat(filepath.Join(targetPath, targFile))

		// Split the file name for reporting
		fileName := strings.SplitN(file.Name(), "-", 2)

		// Report information
		fmt.Println()
		fmt.Println(fmt.Sprintf("Info about %s:\n  Zettlekasten ID: %s\n  Creation Time: %s\n  Last Modified: %s",
			// Name
			fileName[1],
			// ZettkID
			fileName[0],
			// cTime - This is a bad solution as it parses from the ZettK ID
			cTime(fileName[0]),
			// mTime
			file.ModTime().Format("2006-01-02 15:04"))) // Mod Time

		// File opening logic - Fires only if --open is passed
		openFile, err := cmd.Flags().GetBool("open")
		if openFile {
			editor := exec.Command(os.Getenv("EDITOR"), filepath.Join(targetPath, targFile))
			editor.Stdin = os.Stdin
			editor.Stdout = os.Stdout
			editor.Stderr = os.Stderr
			editor.Run()
		}
	},
}

// I hate how these are hard-coded, but probably fine since the date format
// does not change and is auto-generated
func cTime(ctime string) string {
	s := ctime
	yIndex := 4
	mIndex := 6
	dtSplit := 8
	timeIndex := 10
	hrIndex := 12
	formattedTime := s[:yIndex] + "-" + s[yIndex:mIndex] + "-" + s[mIndex:dtSplit] + " " + s[dtSplit:timeIndex] + ":" + s[timeIndex:hrIndex]
	return formattedTime
}

// fileTraversal
// Custom traversal function to add all file names which match a given regexp
// to a slice which is then returned
func fileTraversal(target string, arg string) []string{
	files, err := os.ReadDir(target)
	if err != nil {fmt.Println(err)}
	var foundFiles []string = make([]string, 0, 10)
	targetRegexp := regexp.MustCompile(".*"+arg+".*")
	for _, file := range files {
		isMatch := targetRegexp.MatchString(file.Name())
		if isMatch {
			foundFiles = append(foundFiles, file.Name())
		}
	}
	return foundFiles
}

// findFile
func findFile(matchingFiles []string, arg string) string {
	var ans int
	fmt.Println(fmt.Sprintf("Multiple files matching search term \"%s\" found.",
		arg))
	for i:= 0; i < len(matchingFiles); i++ {
		fmt.Println(fmt.Sprintf("  %d) %s", i, matchingFiles[i]))
	}
	fmt.Println()
	fmt.Print(fmt.Sprintf("Which file do you want? [0-%d] ", len(matchingFiles)-1))
	fmt.Scan(&ans)
	for (ans < 0 || ans >= len(matchingFiles)) {
		fmt.Println("Outside range.")
		fmt.Println()
		fmt.Print(fmt.Sprintf("Which file do you want? [0-%d] ", len(matchingFiles)-1))
		fmt.Scan(&ans)
	}
	return matchingFiles[ans]
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringP("target", "t", "inbox", "Specify the target location")
}
