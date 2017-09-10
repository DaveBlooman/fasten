package fileutils

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/DaveBlooman/fasten/appmeta"
	"github.com/DaveBlooman/fasten/connect"
	"github.com/DaveBlooman/fasten/output"
	"github.com/DaveBlooman/fasten/workflow"
)

func CopyFiles(app appmeta.Application, stack appmeta.AppStack, appDestination string, connection connect.SSHClient) error {
	fileList, dirList, err := getFiles(app.Path)
	if err != nil {
		return err
	}

	for _, dir := range dirList {

		dirName := strings.SplitAfter(dir, app.Path)[1]

		err := connection.MakeDir(appDestination + dirName)
		if err != nil {
			return err
		}
	}

	for _, file := range fileList {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		fileName := strings.SplitAfter(file, app.Path)[1]
		fileCopy := connection.CopyFile(data, fileName, appDestination)
		if fileCopy != nil {
			return fileCopy
		}
	}

	return nil
}

func BuildLocal(localBuild bool, app appmeta.Application) error {
	if !localBuild {
		return nil
	}

	command := "GOOS=linux GOARCH=amd64 go build -o main"
	cmd := exec.Command("sh", "-c", command)
	cmd.Env = os.Environ()
	cmd.Dir = app.Path

	b, err := cmd.CombinedOutput()
	if err != nil {
		output.Error(string(b))
		return err
	}

	return nil

}
func getFiles(filesDir string) ([]string, []string, error) {

	ignoreFile, err := workflow.ReadIgnoreFile(".")
	if err != nil {
		return nil, nil, err

	}

	files, err := workflow.LoadFiles(filesDir, ignoreFile)
	if err != nil {
		return nil, nil, err
	}

	fileList := []string{}
	dirList := []string{}

	err = filepath.Walk(filesDir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if f.IsDir() {
			dirList = append(dirList, path)
			return nil
		}
		fileList = append(fileList, path)
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	var deployableFiles []string

	for _, file := range fileList {
		t := contains(filesDir, files, file)
		if t {
			deployableFiles = append(deployableFiles, file)
		}
	}

	return deployableFiles, dirList, nil
}

func contains(dir string, s []string, e string) bool {
	for _, a := range s {
		if dir+"/"+a == e {
			return true
		}
	}
	return false
}
