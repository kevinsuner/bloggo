// SPDX-License-Identifier: MIT
package generator

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func GetFilesInDir(dir string) ([]string, error) {
	files := make([]string, 0)

	if err := fs.WalkDir(os.DirFS(dir), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		files = append(files, fmt.Sprintf("%s/%s", dir, path))
		return nil
	}); err != nil {
		return nil, err
	}

	return files[1:], nil
}

func GetFileData(path string) (string, []byte, error) {
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

	return info.Name(), filebyte, nil
}

func GetTemplateFiles(theme string) ([]string, error) {
	root := fmt.Sprintf("%s/%s", THEMES_DIR, theme)
	templates := make([]string, 0)

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".html") {
			templates = append(templates, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return templates, nil
}
