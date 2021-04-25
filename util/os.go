package util

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/levibostian/dotenv/ui"
)

func GetFileContents(path string) *string {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil
	}
	if info.IsDir() {
		return nil
	}
	content, err := ioutil.ReadFile(path)
	ui.HandleError(err)

	fileContentString := string(content)

	return &fileContentString
}

func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

// exists returns whether the given file or directory exists
func IsDirExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return stat.IsDir(), nil
}

// Assert we have a diretory and make sure that path ends with 1 `/` character at end.
func SanitizeDirectory(path string) (string, error) {
	if path == "" {
		path = "./"
	}

	var err error
	// convert to abolute path to make debugging easier. knowing an absolute path over relative is easier to debug
	path, err = filepath.Abs(path)
	if err != nil {
		return "", nil
	}

	exists, err := IsDirExists(path)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", errors.New(fmt.Sprintf("Directory at path %s does not exist", path))
	}

	// Last, before we return the valid path, add a trailing / if it does not exist.
	if path[len(path)-1:] != "/" {
		path = path + "/"
	}

	return path, nil
}

func Pwd() string {
	pwd, _ := os.Getwd()
	return fmt.Sprintf("%s/", pwd)
}
