// SPDX-License-Identifier: MIT
package main

import (
	"bloggo/render"
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

const (
	CONTENT_DIR       string = "./content"
	PUBLIC_DIR        string = "./public"
	THEMES_DIR        string = "./themes"
	CONTENT_POSTS_DIR string = CONTENT_DIR + "/posts"
	PUBLIC_POSTS_DIR  string = PUBLIC_DIR + "/posts"
	PUBLIC_ASSETS_DIR string = PUBLIC_DIR + "/assets"
	PUBLIC_CSS_DIR    string = PUBLIC_DIR + "/css"
)

func readYAML() error {
	viper.SetConfigName("bloggo")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}

func main() {
	var baseURL string
	var themeDir string
	var limit int
	var serve bool

	flag.StringVar(&baseURL, "base-url", "http://127.0.0.1:5500/public", "--base-url=http://127.0.0.1:5500/public")
	flag.StringVar(&themeDir, "theme", "bloggo", "--theme=bloggo")
	flag.IntVar(&limit, "limit", 10, "--limit=10")
	flag.BoolVar(&serve, "serve", false, "--serve")
	flag.Parse()

	if err := readYAML(); err != nil {
		log.Fatalf("%s: %v\n", "failed to load yaml", err)
	}

	for _, page := range viper.GetStringSlice("pages") {
		switch page {
		case "index":
			fns := make([]render.RendererOption, 0)
			for _, key := range viper.AllKeys() {
				if fn, ok := render.RendererFunctions[key]; ok {
					if !strings.Contains(key, ".") {
						fns = append(fns, fn)
					}

					if strings.HasPrefix(key, "index") {
						fns = append(fns, fn)
					}
				}
			}

			fmt.Printf("%+v\n", render.New(fns...))
		case "about":
			println("unimplemented!")
		case "archive":
			println("unimplemented!")
		case "404":
			println("unimplemented!")
		default:
			log.Fatalln("not a supported page")
		}
	}

	fmt.Println(viper.GetString("base-url"))
	fmt.Println(viper.GetString("theme"))
	fmt.Println(viper.GetStringSlice("pages"))
	fmt.Println(viper.GetInt("index.posts.limit"))

	var assetsDir string = fmt.Sprintf("%s/%s/%s", THEMES_DIR, themeDir, "assets")
	var cssDir string = fmt.Sprintf("%s/%s/%s", THEMES_DIR, themeDir, "css")

	if err := copyDir(assetsDir, PUBLIC_ASSETS_DIR); err != nil {
		log.Fatalf("%s: %v\n", "failed to copy dir", err)
	}

	if err := copyDir(cssDir, PUBLIC_CSS_DIR); err != nil {
		log.Fatalf("%s: %v\n", "failed to copy dir", err)
	}

	generator := newGenerator(baseURL, themeDir)
	indexErr := generator.index(limit)
	aboutErr := generator.page("about")
	notFoundErr := generator.page("404")
	archiveErr := generator.archive()
	postsErr := generator.posts()
	if err := errors.Join(
		indexErr,
		aboutErr,
		notFoundErr,
		archiveErr,
		postsErr,
	); err != nil {
		log.Fatalf("%s: %v\n", "failed to generate html", err)
	}

	if serve {
		app := fiber.New()
		app.Static("/public", "./public")
		app.Static("/public/assets", "./public/assets")
		app.Static("/public/css", "./public/css")
		app.Static("/public/posts", "./public/posts")
		port, _, _ := strings.Cut(strings.SplitAfterN(baseURL, ":", -1)[2], "/")
		if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
			log.Fatalf("%s: %v", "failed to start server", err)
		}
	}
}
