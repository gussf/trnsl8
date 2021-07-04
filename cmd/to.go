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

	"github.com/trnsl8/internal"
	"github.com/trnsl8/model"

	"github.com/spf13/cobra"
)

var fileToTranslate = ""
var outputFile = ""
var targetLanguage = ""
var result = model.TranslationResult{"", "", ""}
var err = errors.New("")

// toCmd represents the to command
var toCmd = &cobra.Command{
	Use:   "to",
	Short: "Translates an input to a specified language code",
	Long: `
The trnsl8 'to' command expects you to provide a language code and the text you wish to be translated
Examples:
	trnsl8 to en "Quero traduzir esta frase"
	trnsl8 to ja I want this sentence in japanese
	trnsl8 to es -f file.txt -o translation.txt`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println(errors.New("insufficient number of arguments provided to command 'trnsl8 to'"))
			return
		}
		targetLanguage = args[0]

		// --file flag was set, try to open and translate from file
		if fileToTranslate != "" {
			fmt.Println("Translating text from ", fileToTranslate, "...")
			result, err = internal.TranslateToTargetLanguageFromFile(targetLanguage, fileToTranslate)
		} else {
			if len(args) < 2 {
				fmt.Println(errors.New("insufficient number of arguments provided to command 'trnsl8 to'"))
			} else {
				text := strings.Join(args[1:], " ")
				result, err = internal.TranslateToTargetLanguage(targetLanguage, &text)
			}

		}

		if err != nil {
			fmt.Println(err)
			return
		}

		// If flag --output was specified, save output to file
		if outputFile != "" {
			fmt.Println("Writing translation into", outputFile, "...")

			_, err = internal.WriteTranslationToFile(result, outputFile)

			if err != nil {
				fmt.Println(err)
				fmt.Println("\n---------- WARNING: Translation result could NOT be saved to output file ----------")
				fmt.Println(result)
			} else {
				fmt.Println("Success")
			}

		} else {
			fmt.Println(result)
		}
	},
}

func init() {
	rootCmd.AddCommand(toCmd)
	toCmd.PersistentFlags().StringVarP(&fileToTranslate, "file", "f", "", "Translate from a specific file")
	toCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "", "File in which to write the resulting translation")
}
