/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initalize a Zettlekasten",
	Long: `Initalize a new Zettlekasten with subdirectories. The Zettlekasten
is provided in as "vanilla" a form as possible, considering the
logistics of digital databases. Note that if a Zettlekasten exists
at the supplied location, it will ask for permission before overwriting.`,
	Run: func(cmd *cobra.Command, args []string) {
		
		// Grab the home directory and build the directory path
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Failed to find user home directory")
		}
		baseDir := filepath.Join(homeDir, "zettlekasten")

		// Check if a Zettlekasten already exists at the given location
		// If it does exist, prompt the user
		if _, err := os.Stat(baseDir); !os.IsNotExist(err) {
			var ans string = "N"
			fmt.Println("=== WARNING ===")
			fmt.Println("Directory exists at", baseDir)
			fmt.Print("Overwrite existing directory? (Y/N) ")
			fmt.Scan(&ans)
			if err != nil {fmt.Println(err)}
			if strings.ToUpper(ans) != "Y" {
				fmt.Println("Aborting")
				return
			} else {
				os.RemoveAll(baseDir)
			}
   		}

		zettDirs := []string{"00-INBOX", "01-ARCHIVE", "02-REFERENCES", "01-ARCHIVE/daily-notes", "templates"}

		// Create the folder structure
		for _, zettDir := range zettDirs {
			err := os.MkdirAll(filepath.Join(baseDir,zettDir), os.ModePerm)
			if err != nil {fmt.Println(err)}
		}

		// Create note.md default template and write to templates
		noteContents := `---
type: note
title: %s
created: %s
tags:
    - 
---`

		defaultTemp := filepath.Join(baseDir, "templates", "note.md")
		file, err := os.Create(defaultTemp)
		if err != nil {
			fmt.Println("Failed to create markdown file", err)
		}
		defer file.Close()
		_, err = file.WriteString(noteContents)
		if err != nil {
			fmt.Println("Failed to write file")
		}
		file.Close()

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
