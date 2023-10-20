// SPDX-License-Identifier: MIT
package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

const REPO_URL string = "https://github.com/itsksrof/bloggo.git"
const REPO_DIR string = "bloggo-master"

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates core components by pulling changes from master",
	Long: `
Temporarily clones the latest version of Bloggo from it's master branch,
and updates all its core components by replacing the old ones for the
new ones.
	`,
	Run: update,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func update(cmd *cobra.Command, args []string) {
	if err := exec.Command("git", "clone", REPO_URL, REPO_DIR).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	files := []string{
		"bin",
		"commands",
		"generator",
		"themes/bloggo",
		"go.mod",
		"go.sum",
		"LICENSE",
		"main.go",
	}

	for _, file := range files {
		if err := os.RemoveAll(file); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := os.Rename(fmt.Sprintf("%s/%s", REPO_DIR, file), file); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	if err := os.RemoveAll(REPO_DIR); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
