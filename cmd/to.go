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

// toCmd represents the to command
var toCmd = &cobra.Command{
	Use:   "to",
	Short: "Translates an input to a specified language code",
	Long: `
The trnsl8 'to' command expects you to provide a language code and the text you wish to be translated
Examples:
	trnsl8 to en "Quero traduzir esta frase"
	trnsl8 to ja I want this sentence in japanese`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println(errors.New("insufficient number of arguments provided to command 'trnsl8 to'"))
		} else {
			text := strings.Join(args[1:], " ")
			api.TranslateToTargetLanguage(args[0], &text)
		}
	},
}

func init() {
	rootCmd.AddCommand(toCmd)
	toCmd.Flags().BoolP("verbose", "v", false, "Print additional info during translation")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
