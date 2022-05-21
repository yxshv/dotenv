package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var slashT string = "   "

var rootCmd = &cobra.Command{
	Use:   "dotenv",
	Short: "Dotenv allows u to manage your .env files",
	Long:  "Dotenv allows u to manage your .env files",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\n%sWelcome to Dotenv\n\n", slashT)

		fmt.Printf("%sCommands:\n\n", slashT)
		fmt.Printf("%s%slist - List all the variables in the .env file\n", slashT, slashT)
		fmt.Printf("%s%sset <NAME> <VALUE> - Set a variable from the file\n", slashT, slashT)
		fmt.Printf("%s%sremove <NAME> - Removes a variable from the file", slashT, slashT)
		fmt.Println("\n ")

		fmt.Printf("%sFlags:\n", slashT)
		fmt.Printf("\n%s%sfile - The file to read from\n\n", slashT, slashT)
	},
}

func init() {
	rootCmd.PersistentFlags().String("file", "", "The file to read from")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
