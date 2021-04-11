package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/levibostian/dotenv/ui"
)

func GetFileContents(path string, fileDescribe string) []byte {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		ui.Abort(fmt.Sprintf("%s file at path, %s, does not exist", fileDescribe, path))
	}
	if info.IsDir() {
		ui.Abort(fmt.Sprintf("%s file at path, %s, is a directory and not a file.", fileDescribe, path))
	}
	content, err := ioutil.ReadFile(path)
	ui.HandleError(err)

	return content
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
