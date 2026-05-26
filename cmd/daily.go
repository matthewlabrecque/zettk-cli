/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
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

// dailyCmd represents the daily command
var dailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "Open the daily note",
	Long: `Opens the daily note located in the daily-notes directory.
If the daily note doesn't yet exist, it will automatically
create it.`,
	Args: cobra.ExactArgs(0), // Temp fix until we can solve cobra.NoArgs
	Run: func(cmd *cobra.Command, args []string) {
		// Find the Zettlekasten direcory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Failed to find user home directory")
		}
		zettkDir := filepath.Join(homeDir, "zettlekasten")

		// Create the daily note
		dNote := filepath.Join(zettkDir, "daily-notes", fmt.Sprintf("%s.md", time.Now().Format("2006-01-02")))
		dFile, err := os.OpenFile(dNote, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Failed to create daily note")
		}
		defer dFile.Close()
		dFile.Close()

		// Run neovim to open the note
    		nvim := exec.Command("nvim", dNote)
		nvim.Stdin = os.Stdin
		nvim.Stdout = os.Stdout
		nvim.Stderr = os.Stderr
		nvim.Run()

	},
}

func init() {
	rootCmd.AddCommand(dailyCmd)

}
