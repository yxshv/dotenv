package cmd

import (
	"fmt"
	"os"
  "strings"

	"github.com/spf13/cobra"
)

func parseVar(s, sep string) (string, string) {
    x := strings.Split(s, sep)
    if len(x) > 2 {
        name, rest := x[0], strings.Join(x[1:],"")
        return name, rest
    }
    fmt.Println(x)
    return x[0], x[1]
}

func stringify(s map[string]string) string {
    result := ""

    for k,v := range s {
        result += k + "=" + v
    }

    return result
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

var slashT string = "   "

var rootCmd = &cobra.Command{
	Use:   "dotenv",
	Short: "Dotenv allows u to manage your .env files",
	Long:  "Dotenv allows u to manage your .env files",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\n%sWelcome to Dotenv\n\n", slashT)

		fmt.Printf("%sCommands:\n\n", slashT)
		fmt.Printf("%s%slist - List all the variables in the .env file\n", slashT, slashT)
		fmt.Printf("%s%schange <NAME> <VALUE> - Changes the variable if it exists else creates it\n", slashT, slashT)
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
