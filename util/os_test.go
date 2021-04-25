package util

import (
	"testing"

	assert "github.com/go-playground/assert/v2"
)

// SanitizeDirectory

func Test_SanitizeDirectory_givenEmptyString_expectCurrentDirectory(t *testing.T) {
	given := ""
	actual, err := SanitizeDirectory(given)

	assert.Equal(t, err, nil)
	assert.Equal(t, actual, Pwd())
}

func Test_SanitizeDirectory_givenDot_expectCurrentDirectory(t *testing.T) {
	given := "."
	actual, err := SanitizeDirectory(given)

	assert.Equal(t, err, nil)
	assert.Equal(t, actual, Pwd())
}

func Test_SanitizeDirectory_givenDirectoryWithoutDash_expectDirectoryWithDash(t *testing.T) {
	given := "../util"
	actual, err := SanitizeDirectory(given)

	assert.Equal(t, err, nil)
	assert.Equal(t, actual, Pwd())
}

func Test_SanitizeDirectory_givenDirectoryWithDash_expectSameResult(t *testing.T) {
	given := "../util/"
	actual, err := SanitizeDirectory(given)

	assert.Equal(t, err, nil)
	assert.Equal(t, actual, Pwd())
}

func Test_SanitizeDirectory_givenDirectoryThatDoesntExist_expectError(t *testing.T) {
	given := "../doesnotexist/"
	actual, err := SanitizeDirectory(given)

	assert.NotEqual(t, err, nil)
	assert.Equal(t, actual, "")
}

func Test_SanitizeDirectory_givenFile_expectError(t *testing.T) {
	given := "../util/os.go"
	actual, err := SanitizeDirectory(given)

	assert.NotEqual(t, err, nil)
	assert.Equal(t, actual, "")
}
