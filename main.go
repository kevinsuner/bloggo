// SPDX-License-Identifier: MIT
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

const (
	PUBLIC_DIR  string = "public"
	THEMES_DIR  string = "themes"
	CONTENT_DIR string = "content"
	POSTS_DIR   string = "posts"
)

var (
	ErrLoadConfig    error = errors.New("failed to load configuration file")
	ErrCopyFiles     error = errors.New("failed to copy files")
	ErrCreateDir     error = errors.New("failed to create directory")
	ErrKeyNotInMap   error = errors.New("failed to find key in map")
	ErrInitGenerator error = errors.New("failed to initialize generator")
	ErrGenerate      error = errors.New("failed to generate")
	ErrPosts         error = errors.New("failed to get posts")
	ErrHTTPServer    error = errors.New("failed to start http server")
)

func main() {
	viper.SetConfigName("bloggo")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("[viper.ReadInConfig] error: %v\n", ErrLoadConfig)
	}

	err := CopyDirsFromTheme(viper.GetString("theme"), viper.GetStringSlice("dirs"))
	if err != nil {
		log.Fatalf("[CopyDirsFromTheme] error: %v\n", ErrCopyFiles)
	}

	err = os.Mkdir(fmt.Sprintf("%s/%s", PUBLIC_DIR, POSTS_DIR), 0755)
	if err != nil {
		log.Fatalf("[os.Mkdir] error: %v\n", ErrCreateDir)
	}

	for _, page := range viper.GetStringSlice("pages") {
		opts, ok := GeneratorOptions[page]
		if !ok {
			log.Fatalf("[GeneratorOptions] error: %v\n", ErrKeyNotInMap)
		}

		generator, err := NewGenerator(fmt.Sprintf("%s/%s.md", CONTENT_DIR, page), opts...)
		if err != nil {
			log.Fatalf("[NewGenerator] error: %v: %v\n", ErrInitGenerator, err)
		}

		err = generator.Generate(page)
		if err != nil {
			log.Fatalf("[generator.Generate] error: %v: %v\n", ErrGenerate, err)
		}
	}

	files, err := GetFilesInDir(fmt.Sprintf("%s/%s", CONTENT_DIR, POSTS_DIR))
	if err != nil {
		log.Fatalf("[GetFilesInDir] error: %v: %v\n", ErrPosts, err)
	}

	for _, file := range files {
		generator, err := NewGenerator(file, WithMarkdown())
		if err != nil {
			log.Fatalf("[NewGenerator] error: %v: %v\n", ErrInitGenerator, err)
		}

		err = generator.Generate("post")
		if err != nil {
			log.Fatalf("[generator.Generate] error: %v: %v\n", ErrGenerate, err)
		}
	}

	if viper.GetBool("test") {
		app := fiber.New()
		app.Static("/public", "./public")
		app.Static("/public/assets", "./public/assets")
		app.Static("/public/css", "./public/css")
		app.Static("/public/posts", "./public/posts")
		port, _, _ := strings.Cut(strings.SplitAfterN(viper.GetString("base-url"), ":", -1)[2], "/")
		if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
			log.Fatalf("[app.Listen] error: %v: %v\n", ErrHTTPServer, err)
		}
	}
}
