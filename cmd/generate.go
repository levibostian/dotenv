package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/levibostian/dotenv/generate"
	"github.com/levibostian/dotenv/lang"
	"github.com/levibostian/dotenv/types"
	"github.com/levibostian/dotenv/ui"
	"github.com/levibostian/dotenv/util"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate lang file from .env values",
	Long: `Generate lang file from .env values. 

For example:

dotenv generate java [args]`,
}

var packageName string
var sourceCodePath string
var outputFilePath string
var inputLang string // if different from output lang
var dotenvFilePath string

func init() {
	var allLangCommands = []*cobra.Command{}
	var langsRequiringPackageName = []*cobra.Command{}
	for _, lang := range lang.GetLangs() {
		langInfo := lang.GetInfo()
		command := &cobra.Command{
			Use:   langInfo.Name,
			Short: fmt.Sprintf("Generate %s", lang.GetOutputFileName()),
			Run: func(cmd *cobra.Command, args []string) {
				generateForLang(langInfo.Name)
			},
		}

		allLangCommands = append(allLangCommands, command)

		if langInfo.RequirePackageName {
			langsRequiringPackageName = append(langsRequiringPackageName, command)
		}
	}

	cobra.OnInitialize(generateInit)
	rootCmd.AddCommand(generateCmd)

	availableInputLangs := []string{}
	for _, lang := range allLangCommands {
		availableInputLangs = append(availableInputLangs, lang.Use)
	}
	availableInputLangsString := strings.Join(availableInputLangs[:], ", ")

	// All langs have common flags so we add to all.
	for _, langCommand := range allLangCommands {
		langCommand.Flags().StringVarP(&outputFilePath, "output", "o", "", "Output file path (example: build/) (optional. default: current dir)")
		langCommand.Flags().StringVarP(&sourceCodePath, "source", "s", "", "Path where the source code exists (required)")
		langCommand.MarkFlagRequired("source")

		langCommand.Flags().StringVarP(&inputLang, "inputLang", "l", "", fmt.Sprintf("Programming language used by project. Options: %s (optional. default: output lang)", availableInputLangsString))

		langCommand.Flags().StringVarP(&dotenvFilePath, "env", "e", "", ".env file path (example: src/) (optional. default: current dir)")

		generateCmd.AddCommand(langCommand)
	}

	for _, langCommand := range langsRequiringPackageName {
		langCommand.Flags().StringVarP(&packageName, "packageName", "p", "", "Package name (example: com.foo.bar) (required)")
		langCommand.MarkFlagRequired("packageName")
	}
}

func setDefaultValues(defaultInputLang string) {
	if inputLang == "" {
		inputLang = defaultInputLang
	}

	if dotenvFilePath == "" {
		dotenvFilePath = "./"
	}

	if outputFilePath == "" {
		outputFilePath = "./"
	}
}

func generateInit() {
}

func generateForLang(outputLangName string) {
	setDefaultValues(outputLangName)

	var err error
	outputFilePath, err = util.SanitizeDirectory(outputFilePath)
	ui.HandleError(err)
	sourceCodePath, err = util.SanitizeDirectory(sourceCodePath)
	ui.HandleError(err)
	dotenvFilePath, err = util.SanitizeDirectory(dotenvFilePath)
	ui.HandleError(err)

	generate.Execute(types.GenerateOptions{OutputLang: outputLangName, InputLang: inputLang, OutputPath: outputFilePath, SourceCodePath: sourceCodePath, PackageName: packageName, DotenvPath: dotenvFilePath})
}
