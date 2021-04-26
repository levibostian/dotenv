package e2e

import (
	"testing"

	assert "github.com/go-playground/assert/v2"
	run "github.com/levibostian/dotenv/e2e/helpers"
	"github.com/levibostian/dotenv/util"
)

func Test_GenerateJavaFromKotlinSource_expectFileGenerated(t *testing.T) {
	output := run.Command("../../dotenv", "generate java --packageName earth.levi.dotenv --source . --inputLang kotlin")

	assert.Equal(t, output, "")

	expected := util.GetFileContents("./Expected.java")
	actual := util.GetFileContents("./Env.java")

	assert.Equal(t, expected, actual)
}
