/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	name   string
	age    int
	format string
)

// formatCmd represents the format command
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Format user data in text or JSON format",
	Long:  `This command takes user name and age, and outputs the result either in plain text or JSON format, based on the --format flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		user := User{
			Name: name,
			Age:  age,
		}

		switch format {
		case "json":
			output, err := json.MarshalIndent(user, "", " ")
			if err != nil {
				fmt.Println("erorr marshaling JSON:", err)
				os.Exit(1)
			}
			fmt.Println(string(output))
		case "text":
			fmt.Printf("User Name: %s\nUser Age: %d\n", user.Name, user.Age)
		default:
			fmt.Println("Unknown format. use 'json' or 'text'.")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)

	formatCmd.Flags().StringVarP(&name, "name", "n", "", "Your name")
	formatCmd.Flags().IntVarP(&age, "age", "a", 0, "Your age")
	formatCmd.Flags().StringVarP(&format, "format", "f", "text", "Output format: json or text")

	formatCmd.MarkFlagRequired("name")
	formatCmd.MarkFlagRequired("age")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// formatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// formatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
