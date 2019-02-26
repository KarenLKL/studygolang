package file

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type fileHandlError string

func (f fileHandlError) Error() string {
	return f.Message()
}

func (f fileHandlError) Message() string {
	return string(f)
}

const prefix = "/list/"

func Handler(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return fileHandlError("path must start with '" + prefix + "'")
	}
	filePath := request.URL.Path[len("/list/"):]
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	_, err = writer.Write(content)
	if err != nil {
		return err
	}
	return nil
}
