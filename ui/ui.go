package ui

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	cliConfig "github.com/levibostian/dotenv/cliconfig"
)

func ShouldNotHappen(err error) {
	color.Red("[BUG] Something happened that should not have. That means there is probably a bug inside of dotenv.")
	color.Red("Report an issue here: https://github.com/levibostian/dotenv/issues/new, and give this message:")
	fmt.Println(err)
	panic("Exiting...")
}

// HandleError pass in error and we will handle it.
func HandleError(err error) {
	if err != nil {
		Error("\nError encountered!")
		fmt.Println(err)
		os.Exit(1)
	}
}

// Debug - Allows you to put anything you want inside. String, struct, etc. We will print that to the console.
// Help: https://gobyexample.com/string-formatting
func Debug(format string, args ...interface{}) {
	if cliConfig.CliConfig.Debug {
		msg := fmt.Sprintf(format, args...)
		color.Cyan("[DEBUG] " + msg)
	}
}

// Verbose - Allows you to put anything you want inside. String, struct, etc. We will print that to the console.
func Verbose(format string, args ...interface{}) {
	if cliConfig.CliConfig.Verbose || cliConfig.CliConfig.Debug {
		msg := fmt.Sprintf(format, args...)
		color.White(msg)
	}
}

func Abort(message string) {
	Error(message)
	os.Exit(1)
}

// Error show a message in red
func Error(message string) {
	color.Red(message)
}
