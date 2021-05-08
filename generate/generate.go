package generate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	mapset "github.com/deckarep/golang-set"
	"github.com/joho/godotenv"
	"github.com/levibostian/dotenv/lang"
	"github.com/levibostian/dotenv/types"
	"github.com/levibostian/dotenv/ui"
	"github.com/levibostian/dotenv/util"
)

func Execute(options types.GenerateOptions) {
	ui.Debug(`Generating. Options: %+v`, options)

	var err error
	inputLangs, err := lang.GetLangs(options.InputLangs)
	ui.HandleError(err)
	outputLang, err := lang.GetLang(options.OutputLang)

	envValuesFound := mapset.NewSet()

	filepath.Walk(options.SourceCodePath, func(path string, info os.FileInfo, err error) error {
		fileContents := util.GetFileContents(path)
		if fileContents == "" {
			return nil
		}

		var inputLangForFile lang.Lang
		for _, lang := range inputLangs {
			if lang.IsFilenameValid(info.Name()) {
				inputLangForFile = lang
			}
		}
		if inputLangForFile == nil {
			ui.Debug(`File not of any input language, skipping: %s`, path)
			return nil
		}

		ui.Debug(`Checking file: %s`, path)

		for _, line := range strings.Split(fileContents, "\n") {
			ui.Debug("Line of file: %s", line)

			for _, envFound := range inputLangForFile.ParseSourceCodeLine(line) {
				envValuesFound.Add(envFound)
			}
		}

		return nil
	})

	ui.Debug("Set values: %s", envValuesFound.String())

	err = godotenv.Load(options.DotenvPath + ".env")
	ui.HandleError(err)

	var envValues = make(map[string]string)
	envValuesFound.Each(func(elem interface{}) bool {
		envSourceCodeValue := elem.(string)

		value := os.Getenv(strings.ToUpper(envSourceCodeValue))

		if value != "" {
			envValues[envSourceCodeValue] = value
		}

		return false
	})

	ui.Debug("Mapped values: %s", envValues)

	outputFile := options.OutputPath + outputLang.GetOutputFileName()
	util.WriteToFile(outputFile, outputLang.GetOutputFile(envValues, options))
	ui.Verbose(fmt.Sprintf("Wrote file to: %s", outputFile))
}
