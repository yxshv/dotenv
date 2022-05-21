package cmd

import (
	  "github.com/spf13/cobra"
    "os"
    "fmt"
    "strings"
)

var addCmd = &cobra.Command{
	  Use:   "change",
	  Short: "Changes the variable if it exists else creates it",
	  Long:  "Changes the variable if it exists else creates it",
    Run: func(cmd *cobra.Command, args []string) {

        if len(args) < 1 {
            fmt.Println("Your are missing the <NAME> argument.")
            return
        } else if len(args) < 2  {
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

        fmt.Printf("Checking for `%s`\n",fileFlag)

        file, err := os.Open(fileFlag)
        
        if err != nil {
            if strings.HasSuffix(err.Error(), "The system cannot find the file specified.") == true {
                fmt.Println("file not found. so creating one..")

                file, err = os.Create(fileFlag)

                check(err)

                fmt.Println("successfully created")
            }
        }
        
        defer file.Close()

        var content []byte

        file.Read(content)

        if err != nil {
            fmt.Printf("Error occured while reading the file\nError : %s \n", err)
            return
        }
        
        variables := make(map[string]string)

        for _, v := range strings.Split(string(content), "\n") {
            fmt.Println(v)
            name, value := parseVar(v, "=")
            variables[name] = value
        }

        variables[givenName] = givenValue

        content_bytes := []byte(stringify(variables))
        os.WriteFile(fileFlag, content_bytes, 0644)
        check(err)

        fmt.Printf("Successfully add a variable to `%s`\n%s=%s\n",fileFlag,givenName,givenValue)

	  },
}

func init() {
	  rootCmd.AddCommand(addCmd)
}
