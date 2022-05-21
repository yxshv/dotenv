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
		name, rest := x[0], strings.Join(x[1:], "")
		return name, rest
	}
	return x[0], x[1]
}

func stringify(s map[string]string) string {
	result := ""

	for k, v := range s {
		result += k + "=" + v + "\n"
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
	Short: "Dotenv allows you to manage your .env files",
	Long:  "Dotenv allows you to manage your .env files",
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
