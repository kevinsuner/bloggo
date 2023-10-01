package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func getFilesInDir(dir string) ([]string, error) {
	files := []string{}

	if err := fs.WalkDir(os.DirFS(dir), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		files = append(files, fmt.Sprintf("./%s/%s", dir, path))
		return nil
	}); err != nil {
		return nil, err
	}

	return files[1:], nil
}

func getFileData(path string) (string, []byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	filebyte, err := os.ReadFile(path)
	if err != nil {
		return "", nil, err
	}

	info, err := file.Stat()
	if err != nil {
		return "", nil, err
	}

	return strings.TrimRight(info.Name(), ".md"), filebyte, nil
}
