package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a variable from the file",
	Long:  "Removes a variable from the file",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("Your are missing the <NAME> argument.")
			return
		}

		givenName := args[0]

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

		content, err := os.ReadFile(fileFlag)

		if err != nil {
			if strings.HasSuffix(err.Error(), "no such file or directory") == true {
				fmt.Println("file not found. so creating one..")

				os.Create(fileFlag)

				content, err = os.ReadFile(fileFlag)

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

		delete(variables, givenName)

		content_bytes := []byte(stringify(variables))
		os.WriteFile(fileFlag, content_bytes, 0644)
		check(err)

		fmt.Printf("\nSuccessfully removed the variable `%s` (if it existed) from `%s`\n", givenName, fileFlag)

	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
