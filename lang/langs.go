package lang

import (
	"errors"
	"fmt"

	"github.com/levibostian/dotenv/types"
)

/**
Adding a new programming language to the CLI? Great!!

In this file, you will find comments that start with `NL:` (stands for "new language"). Follow the instructions there.
*/

// Information about the programming language. This is used by the CLI code to generate the CLI commands but is also used by each language's code generation file. Let's say your programming language requires that you enter a package name or module name at the top of the source code file. Then, add that new struct property below.
type LangInfo struct {
	// name of the language in lowercase and no spaces. If your programming language is Typescript, for example, return "typescript"
	Name string
	// Does your language require package names? (kotlin, java does)
	RequirePackageName bool
}

type Lang interface {
	GetInfo() LangInfo
	// given a filename, determine if it's a valid filename for your programming language. If given `name.xml`, return false for Typescript. `.tsx` or `.ts` return true for Typescript.
	IsFilenameValid(filename string) bool
	// Given 1 line of code of your language, parse that to get all of the environment variables in that line. Example: For Kotlin if the line is `val foo = Env.bar + Env.fooBar + Env.foo` return ["BAR", "FOO_BAR", "BAR"]. Don't worry about duplicates those are removed for you.
	ParseSourceCodeLine(line string) []string
	// Filename to save the output code in. Example: `Env.java` for Java. `env.ts` for Typescript.
	GetOutputFileName() string
	// Generate a multi-line string that is the output source code with all of the environment variables inside. It's recommended you view the existing languages supported by this code base for examples as there are a lot of steps in this function.
	GetOutputFile(values map[string]string, options types.GenerateOptions) string
}

// NL: Add a case statement for your language.
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
	/**

	NL: Here is some example code you can copy/paste

	case <NameOfLang>{}.GetInfo().Name:
		selectedLang = <NameOfLang>{}
		break
	}

	*/
	if selectedLang == nil {
		return nil, errors.New(fmt.Sprintf("Invalid lang. Provided %s", langString))
	}

	return selectedLang, nil
}

// NL: Add your new language below to the list. If your new language struct is named "Typescript", for example, add `Typescript{}` to the end of this list below.
func GetLangs() []Lang {
	return []Lang{Java{}, Kotlin{}}
}
