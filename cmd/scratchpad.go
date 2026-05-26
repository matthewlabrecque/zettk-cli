/*
Copyright © 2026 Matthew Labrecque <mlabrecque2002@gmail.com> 
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"github.com/spf13/cobra"
)

// scratchpadCmd represents the scratchpad command
var scratchpadCmd = &cobra.Command{
	Use:   "scratchpad",
	Short: "Open the scratchpad note for quick notes",
	Long: `Open a dedicated scratchpad file for jotting down quick notes.
If the scratchpad doesn't exist, it will automatically create it in
the scratchpad directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Find the Zettlekasten direcory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Failed to find user home directory")
		}
		zettkDir := filepath.Join(homeDir, "zettlekasten")

		// Add the new note to the daily note
		spPath := filepath.Join(zettkDir, "scratchpad", "scratchpad.md")
		spFile, err := os.OpenFile(spPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Failed to find or create the scratchpad")
		}
		spFile.Close()

		// Run neovim to open the note
    		nvim := exec.Command("nvim", spPath)
		nvim.Stdin = os.Stdin
		nvim.Stdout = os.Stdout
		nvim.Stderr = os.Stderr
		nvim.Run()
	},
}

func init() {
	rootCmd.AddCommand(scratchpadCmd)

}
