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
		
		// Get the template
		tVal, err := cmd.Flags().GetString("template")
		if err != nil { fmt.Println(err) }
		template, err := os.ReadFile(filepath.Join(zettkDir, "templates", tVal))
		if err != nil { fmt.Println(err) }

		// Create the markdown file
		fName := filepath.Join(zettkDir, fmt.Sprintf("%s.md", filepath.Clean(args[0]))) // Expected value is "path/to/zettk/my-name.md"
		file, err := os.Create(fName)
		if err != nil {
			fmt.Println("Failed to create markdown file", err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf(string(template), filepath.Clean(args[0]), time.Now().Format("2006-01-02")))
		if err != nil {
			fmt.Println("Failed to write file")
		}
		file.Close()

		// Add the new note to the daily note
		dNote := filepath.Join(zettkDir, "daily-notes", fmt.Sprintf("%s.md", time.Now().Format("2006-01-02")))
		dFile, err := os.OpenFile(dNote, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Failed to create daily note")
		}
		defer dFile.Close()
		link := "[[" + fmt.Sprintf("%s.md", filepath.Clean(args[0])) + "]]"
		_, err = dFile.WriteString(link)
		if err != nil {
			fmt.Println(err)
		}
		dFile.Close()

		// Run neovim to open the note
    		nvim := exec.Command("nvim", fName)
		nvim.Stdin = os.Stdin
		nvim.Stdout = os.Stdout
		nvim.Stderr = os.Stderr
		nvim.Run()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Flag to allow custom template spec
	newCmd.Flags().StringP("template", "t", "note.md", "Specify custom note template")
}
