/*
Copyright © 2026 Matthew Labrecque <mlabrecque2002@gmail.com	>

*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a file in the default editor",
	Long: `Open a file in your text editor based on what you search.
If multiple results appear, it will prompt for which
file you want to open.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Failed to find user home directory")
		}

		searchQuerry := filepath.Clean(args[0])

		directories := []string{"INBOX", "ARCHIVE", "INPUT", "PROJECTS", "PRIVATE"}

		// Walk through all the directories
		var foundFiles []string = make([]string, 0, 10)
		targetRegexp := regexp.MustCompile(".*" + searchQuerry + ".*")
		for _, dir := range directories {
			files, err := os.ReadDir(filepath.Join(homeDir, "zettelkasten", dir))
			if err != nil {fmt.Println(err)}
			for _, file := range files {
				if !file.IsDir() && targetRegexp.MatchString(file.Name()) {
					// Note we have to add the dir for file path construction
					// later on
					foundFiles = append(foundFiles, dir + "/" + file.Name())
				}
			}
		}

		// Determine if additional querry is needed
		var foundFile string = ""
		if len(foundFiles) == 0 {
			fmt.Println("Found 0 notes matching search querry:", searchQuerry)
			return
		} else if len(foundFiles) > 1 {
			var ans int
			fmt.Println(fmt.Sprintf("Multiple files matching search term \"%s\" found.", searchQuerry))
			for i:= 0; i < len(foundFiles); i++ {
				fmt.Println(fmt.Sprintf("  %d) %s", i, foundFiles[i]))
			}
			fmt.Println()
			fmt.Print(fmt.Sprintf("Which file do you want? [0-%d] ", len(foundFiles)-1))
			fmt.Scan(&ans)
			for (ans < 0 || ans >= len(foundFiles)) {
				fmt.Println("Please select from within the valid options.")
				fmt.Println()
				fmt.Print(fmt.Sprintf("Which file do you want? [0-%d] ", len(foundFiles)-1))
				fmt.Scanf("%d", &ans)
			}
			foundFile = foundFiles[ans]
		} else {
			foundFile = foundFiles[0]
		}

		// Reconstruct the file path
		fullPath := filepath.Join(homeDir, "zettelkasten", foundFile)

		// Run neovim to open the note
    	editor := exec.Command(os.Getenv("EDITOR"), fullPath)
		editor.Stdin = os.Stdin
		editor.Stdout = os.Stdout
		editor.Stderr = os.Stderr
		editor.Run()
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
