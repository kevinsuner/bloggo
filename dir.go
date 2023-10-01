package main

import (
	"fmt"
	"os"
)

func copyDir(src, dest string) error {
	files, err := getFilesInDir(src)
	if err != nil {
		return err
	}

	err = os.RemoveAll(dest)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	err = os.Mkdir(dest, 0755)
	if err != nil {
		return err
	}

	for _, file := range files {
		filename, filebyte, err := getFileData(file)
		if err != nil {
			return err
		}

		if err := os.WriteFile(
			fmt.Sprintf("%s/%s", dest, filename),
			filebyte,
			0644,
		); err != nil {
			return err
		}
	}

	return nil
}
