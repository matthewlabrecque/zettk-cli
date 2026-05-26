/*
Copyright © 2026 Matthew Labrecque <mlabrecque2002@gmail.com> 
*/
package cmd

import (
	"fmt"
	"time"
	"os"
	"os/exec"
	"path/filepath"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [title]",
	Short: "Create a new note and update the zettlekasten.",
	Long: `Creates a new markdown note and adds it to the daily note for the
current day. If the daily note doesn't yet exist, it creates the daily
note. This command requires exactly one argument - the name of the file.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Find the Zettlekasten direcory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Failed to find user home directory")
		}
		zettkDir := filepath.Join(homeDir, "zettlekasten")
		
		// Create the markdown file
		fName := filepath.Join(zettkDir, fmt.Sprintf("%s.md", filepath.Clean(args[0]))) // Expected value is "path/to/zettk/my-name.md"
		file, err := os.Create(fName)
		if err != nil {
			fmt.Println("Failed to create markdown file", err)
		}
		defer file.Close()

		// See if the daily note exists
		// If not, create the daily note
		dailyNote := filepath.Join(zettkDir, "daily-notes", fmt.Sprintf("%s.md", time.Now().Format("2006-01-02")))

		// TODO: Implement the creation of the daily note

		// Inject the markdown into the daily note
		// TODO: Find a markdown library for this

		// Run neovim to open the note
		exec.Command("nvim", fName)

		// Use this for testing
		fmt.Println("File successfully created!")
		fmt.Println(dailyNote)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.
	newCmd.Flags().StringP("template", "t", "standard.md", "Specify custom note template")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
