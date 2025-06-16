/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var reminderDuration time.Duration

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random [file]",
	Short: "Read a file and remind after a certain time",
	Long:  `This command reads a text file and reminds you if you're still reading after a given time duration.`,
	Args:  cobra.ExactArgs(1), // Require exactly 1 argument (the file name)
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]

		// Start reminder goroutine
		if reminderDuration > 0 {
			go func() {
				time.Sleep(reminderDuration)
				fmt.Printf("\n⏰ You've been reading for %v. Take a break!\n", reminderDuration)
			}()
		}

		// Read and print file
		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("❌ Could not read file: %v\n", err)
			return
		}
		fmt.Println("📄 File content:")
		fmt.Println(string(data))

		// Keep the program running until user presses enter
		fmt.Println("\nPress Enter when you're done reading...")
		fmt.Scanln()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Add the reminder flag (e.g. --reminder=2m)
	randomCmd.Flags().DurationVarP(&reminderDuration, "reminder", "r", 0, "Duration after which to show a reminder (e.g. 2m, 30s)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
