// SPDX-License-Identifier: MIT
package commands

import (
	"bloggo/generator"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Converts Markdown content to HTML using Go templates",
	Long: `
Converts the Markdown files from the content folder to HTML using the user defined
configuration, the provided frontmatter in the files itself, and the power of 
Go templates.
	`,
	Run: generate,
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func generate(cmd *cobra.Command, args []string) {
	var (
		theme string   = viper.GetString("theme")
		dirs  []string = viper.GetStringSlice("dirs")
		pages []string = viper.GetStringSlice("pages")
	)

	if err := generator.CopyDirsFromTheme(theme, dirs); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := os.Mkdir(
		fmt.Sprintf("%s/%s", generator.PUBLIC_DIR, generator.POSTS_DIR),
		0755,
	); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, page := range pages {
		opts, ok := generator.GeneratorOptions[page]
		if !ok {
			fmt.Fprintf(os.Stderr, "%s is not supported\n", page)
			os.Exit(1)
		}

		generator, err := generator.NewGenerator(
			fmt.Sprintf("%s/%s.md", generator.CONTENT_DIR, page), opts...,
		)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := generator.Generate(page); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	files, err := generator.GetFilesInDir(
		fmt.Sprintf("%s/%s", generator.CONTENT_DIR, generator.POSTS_DIR),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, file := range files {
		generator, err := generator.NewGenerator(file, generator.WithMarkdown())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := generator.Generate("post"); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
