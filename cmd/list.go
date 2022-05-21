package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the variables in the .env file",
	Long:  "List all the variables in the .env file",
	Run: func(cmd *cobra.Command, args []string) {
		fileFlag, err := cmd.Flags().GetString("file")

		check(err)

		if fileFlag == "" {
			fileFlag = ".env"
		}

		if strings.HasPrefix(fileFlag, ".env") == false {
			fmt.Println("The file needs to startwith `.env`")
			return
		}

		fmt.Printf("Checking for `%s`\n", fileFlag)

		content, err := ioutil.ReadFile(fileFlag)

		if err != nil {
			if strings.HasSuffix(err.Error(), "no such file or directory") == true {
				fmt.Println("file not found. so creating one..")

				os.Create(fileFlag)

				content, err = ioutil.ReadFile(fileFlag)

				check(err)

				fmt.Println("successfully created")
			}
		}

		variables := make(map[string]string)

		for _, v := range strings.Split(string(content), "\n") {
			if len(v) < 0 || v == "" {
				continue
			}
			name, value := parseVar(v, "=")
			variables[name] = value
		}

		tw := table.NewWriter()

		tw.AppendHeader(table.Row{"Name", "Value"})

		for key, value := range variables {
			tw.AppendRow([]interface{}{key, value})
		}

		fmt.Printf("\n%s\n", tw.Render())

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
