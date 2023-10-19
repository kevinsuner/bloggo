// SPDX-License-Identifier: MIT
package generator

import (
	"fmt"
	"os"
)

func CopyDirsFromTheme(theme string, dirs []string) error {
	if err := os.RemoveAll(PUBLIC_DIR); err != nil {
		return err
	}

	if err := os.Mkdir(PUBLIC_DIR, 0755); err != nil {
		return err
	}

	for _, dir := range dirs {
		files, err := GetFilesInDir(fmt.Sprintf("%s/%s/%s", THEMES_DIR, theme, dir))
		if err != nil {
			return err
		}

		if err := os.Mkdir(fmt.Sprintf("%s/%s", PUBLIC_DIR, dir), 0755); err != nil {
			return err
		}

		for _, file := range files {
			filename, filebyte, err := GetFileData(file)
			if err != nil {
				return err
			}

			if err := os.WriteFile(
				fmt.Sprintf("%s/%s/%s", PUBLIC_DIR, dir, filename),
				filebyte,
				0644,
			); err != nil {
				return err
			}
		}
	}

	return nil
}
