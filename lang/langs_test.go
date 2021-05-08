package lang

import (
	"testing"

	assert "github.com/go-playground/assert/v2"
)

// GetLangs

func Test_GetLangs_givenSingleLanguage_expectListWithOneItem(t *testing.T) {
	given := "java"
	actual, err := GetLangs(given)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(actual), 1)
	assert.Equal(t, actual[0].GetInfo().Name, "java")
}

func Test_GetLangs_givenStringArrayNoSpaces_expectList(t *testing.T) {
	given := "java,kotlin"
	actual, err := GetLangs(given)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(actual), 2)
	assert.Equal(t, actual[0].GetInfo().Name, "java")
	assert.Equal(t, actual[1].GetInfo().Name, "kotlin")
}

func Test_GetLangs_givenStringArraySpaces_expectList(t *testing.T) {
	given := "java, kotlin"
	actual, err := GetLangs(given)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(actual), 2)
	assert.Equal(t, actual[0].GetInfo().Name, "java")
	assert.Equal(t, actual[1].GetInfo().Name, "kotlin")
}

func Test_GetLangs_givenStringArray_givenOneLanguageNotValid_expectError(t *testing.T) {
	given := "java, not-language"
	actual, err := GetLangs(given)

	assert.NotEqual(t, err, nil)
	assert.Equal(t, actual, nil)
}

// GetLang

func Test_GetLang_givenLanguageString_expectGetLang(t *testing.T) {
	given := "java"
	actual, err := GetLang(given)

	assert.Equal(t, err, nil)
	assert.Equal(t, actual.GetInfo().Name, "java")
}

func Test_GetLang_givenInvalidLanguage_expectError(t *testing.T) {
	given := "not-language"
	actual, err := GetLang(given)

	assert.NotEqual(t, err, nil)
	assert.Equal(t, actual, nil)
}
