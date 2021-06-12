/*
Copyright © 2021 GUSSF

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"strings"

	"internal/api"

	"github.com/spf13/cobra"
)

var fileToDetect = ""

// detectCmd represents the detect command
var detectCmd = &cobra.Command{
	Use:   "detect",
	Short: "Detects the dominant language in an input",
	Long: `The trnsl8 'detect' command tries to determine the prevalent language in an input

Examples:
	trnsl8 detect "This is in english"`,

	Run: func(cmd *cobra.Command, args []string) {

		// --file flag was set, try to open and detect the text in the filepath
		if fileToDetect != "" {
			fmt.Println("Detecting most prevalent language in", fileToDetect, "...")
			result, err := api.DetectDominantLanguageInFile(fileToDetect)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(result)
			return
		}

		// --file flag was not set, detect text input from args
		if len(args) == 0 {
			fmt.Println(errors.New("insufficient number of arguments provided to command 'trnsl8 detect'"))
		} else {
			input := strings.Join(args, " ")
			result, err := api.DetectDominantLanguageIn(&input)
			fmt.Println(result, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(detectCmd)
	detectCmd.PersistentFlags().StringVarP(&fileToDetect, "file", "f", "", "Detect language in a specific file")
}
