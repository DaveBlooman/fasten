package workflow

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/rliebling/gitignorer"
)

func ReadIgnoreFile(dir string) ([]byte, error) {
	path := filepath.Join(dir, ".fastenignore")

	b, err := ioutil.ReadFile(path)

	if os.IsNotExist(err) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return b, nil
}

func LoadFiles(root string, ignoreFile []byte) (files []string, err error) {
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		mode := info.Mode()
		if !(mode.IsDir() || mode.IsRegular() || mode&os.ModeSymlink == os.ModeSymlink) {
			return nil
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		matched, err := gitignorer.GitIgnore(bytes.NewReader(ignoreFile), rel)
		if err != nil {
			return err
		}

		if mode.IsDir() && matched {
			return filepath.SkipDir
		} else if mode.IsDir() || matched {
			return nil
		}

		files = append(files, rel)

		return nil
	})

	return
}
