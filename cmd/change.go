package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "change",
	Short: "Changes the variable if it exists else creates it",
	Long:  "Changes the variable if it exists else creates it",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("Your are missing the <NAME> argument.")
			return
		} else if len(args) < 2 {
			fmt.Println("Your are missing the <VALUE> argument.")
			return
		}

		givenName, givenValue := args[0], args[1]

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

		variables[givenName] = givenValue

		content_bytes := []byte(stringify(variables))
		os.WriteFile(fileFlag, content_bytes, 0644)

		fmt.Printf("\nSuccessfully add a variable to `%s`\n\n%s=%s\n", fileFlag, givenName, givenValue)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
