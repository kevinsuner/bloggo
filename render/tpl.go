// SPDX-License-Identifier: MIT
package render

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

const THEMES_DIR = "themes"

func getTemplateFiles(theme string) ([]string, error) {
	var root string = fmt.Sprintf("%s/%s", THEMES_DIR, theme)
	var templates []string = make([]string, 0)

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
