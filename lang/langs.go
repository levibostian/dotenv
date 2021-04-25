package lang

import (
	"errors"
	"fmt"

	"github.com/levibostian/dotenv/types"
)

type LangInfo struct {
	Name               string
	RequirePackageName bool
}

type Lang interface {
	GetInfo() LangInfo
	IsFilenameValid(filename string) bool
	ParseSourceCodeLine(line string) []string
	GetOutputFileName() string
	GetOutputFile(values map[string]string, options types.GenerateOptions) string
}

func GetLang(langString string) (Lang, error) {
	var selectedLang Lang

	switch langString {
	case Java{}.GetInfo().Name:
		selectedLang = Java{}
		break
	case Kotlin{}.GetInfo().Name:
		selectedLang = Kotlin{}
		break
	}
	if selectedLang == nil {
		return nil, errors.New(fmt.Sprintf("Invalid lang. Provided %s", langString))
	}

	return selectedLang, nil
}

func GetLangs() []Lang {
	return []Lang{Java{}, Kotlin{}}
}
