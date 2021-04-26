package e2e

import (
	"strings"
	"testing"

	assert "github.com/go-playground/assert/v2"
	run "github.com/levibostian/dotenv/e2e/helpers"
)

func Test_Version_expectValidValue(t *testing.T) {
	output := run.Command("../dotenv", "version")

	assert.NotEqual(t, output, "")
	assert.Equal(t, strings.Count(output, "."), 2)
}
