package generate

import (
	"fmt"

	"github.com/levibostian/dotenv/ui"
	"github.com/levibostian/dotenv/util"
)

func Execute(lang string, outputPath string) {
	ui.Debug(lang)
	ui.Debug(`Output path %s`, outputPath)

	util.WriteToFile(outputPath+`Env.kt`, fmt.Sprintf(kotlinFileTemplate, "com.foo.bar", "val bar: String = \"\""))
	// TODO iterate source code files that match a file pattern (example: <given-input-path>/**/*.kt)
	// TODO using regex, find Env values
	// TODO output to file.

	return
}
