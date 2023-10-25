// SPDX-License-Identifier: MIT
package generator

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/viper"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

const (
	PUBLIC_DIR  string = "public"
	THEMES_DIR  string = "themes"
	CONTENT_DIR string = "content"
	POSTS_DIR   string = "posts"
)

var GeneratorOptions map[string][]GeneratorOption = map[string][]GeneratorOption{
	"index":   {WithMarkdown(), WithPostsIndex()},
	"about":   {WithMarkdown()},
	"archive": {WithMarkdown(), WithPostsArchive()},
	"404":     {WithMarkdown()},
}

type GeneratorOption func(*Generator) error

type Generator struct {
	File       string
	Filename   string
	Lang       string
	BaseURL    string
	Theme      string
	PostsLimit int

	Content template.HTML
	Meta    *Meta
	Posts   []Post
	Archive []Archive
}

type Meta struct {
	Title       string
	RawTitle    string
	Description string
	Author      string
	Robots      string
	Type        string
	Section     string
	Published   string
	Modified    string
	Keywords    string
}

type Post struct {
	ID          int
	Title       string
	Description string
	Author      string
	Published   string
	Keywords    string
	URL         template.URL
}

type Archive struct {
	ID    int
	Posts []Post
}

func NewGenerator(file string, opts ...GeneratorOption) (*Generator, error) {
	generator := &Generator{
		File:       file,
		Lang:       viper.GetString("lang"),
		BaseURL:    viper.GetString("base-url"),
		Theme:      viper.GetString("theme"),
		PostsLimit: viper.GetInt("posts-limit"),
	}

	for _, opt := range opts {
		if err := opt(generator); err != nil {
			return nil, err
		}
	}

	return generator, nil
}

func (g *Generator) Generate(name string) error {
	templates, err := GetTemplateFiles(g.Theme)
	if err != nil {
		return err
	}

	tpl, err := template.New(name).ParseFiles(templates...)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, g); err != nil {
		return err
	}

	if g.Meta.Type == "article" {
		if err := os.WriteFile(
			fmt.Sprintf("%s/%s/%s.html", PUBLIC_DIR, POSTS_DIR, g.Filename),
			buf.Bytes(),
			0644,
		); err != nil {
			return err
		}

		return nil
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/%s.html", PUBLIC_DIR, g.Filename),
		buf.Bytes(),
		0644,
	); err != nil {
		return err
	}

	return nil
}

func WithMarkdown() GeneratorOption {
	return func(g *Generator) error {
		filename, filebyte, err := GetFileData(g.File)
		if err != nil {
			return err
		}
		g.Filename = strings.TrimSuffix(filename, ".md")

		var buf bytes.Buffer
		ctx := parser.NewContext()
		if err := goldmark.New(goldmark.WithExtensions(extension.GFM, meta.Meta)).
			Convert(filebyte, &buf, parser.WithContext(ctx)); err != nil {
			return err
		}
		g.Content = template.HTML(buf.String())

		metadata := meta.Get(ctx)
		meta := new(Meta)

		if metadata["type"].(string) == "article" {
			meta.RawTitle = metadata["raw-title"].(string)
			meta.Section = metadata["section"].(string)
			meta.Published = metadata["published"].(string)
			meta.Modified = metadata["modified"].(string)
		}

		meta.Title = metadata["title"].(string)
		meta.Description = metadata["description"].(string)
		meta.Author = metadata["author"].(string)
		meta.Robots = metadata["robots"].(string)
		meta.Type = metadata["type"].(string)
		meta.Keywords = metadata["keywords"].(string)

		g.Meta = meta
		return nil
	}
}

func WithPostsIndex() GeneratorOption {
	return func(g *Generator) error {
		files, err := GetFilesInDir(fmt.Sprintf("%s/%s", CONTENT_DIR, POSTS_DIR))
		if err != nil {
			return err
		}

		posts := make([]Post, 0)
		for _, file := range files {
			filename, filebyte, err := GetFileData(file)
			if err != nil {
				return err
			}

			markdown := goldmark.New(
				goldmark.WithExtensions(
					meta.Meta,
				),
			)

			var buf bytes.Buffer
			ctx := parser.NewContext()
			if err := markdown.Convert(filebyte, &buf, parser.WithContext(ctx)); err != nil {
				return err
			}

			metadata := meta.Get(ctx)
			id, err := strconv.Atoi(strings.ReplaceAll(metadata["published"].(string), "-", ""))
			if err != nil {
				return err
			}

			posts = append(posts, Post{
				ID:          id,
				Title:       metadata["raw-title"].(string),
				Description: metadata["description"].(string),
				Author:      metadata["author"].(string),
				Published:   metadata["published"].(string),
				Keywords:    metadata["keywords"].(string),
				URL: template.URL(
					fmt.Sprintf(
						"%s/%s/%s.html",
						g.BaseURL,
						"posts",
						strings.TrimSuffix(filename, ".md"),
					),
				),
			})
		}

		sort.SliceStable(posts, func(i, j int) bool {
			return posts[j].ID < posts[i].ID
		})

		if len(posts) >= g.PostsLimit {
			g.Posts = posts[:g.PostsLimit]
			return nil
		}

		g.Posts = posts
		return nil
	}
}

func WithPostsArchive() GeneratorOption {
	return func(g *Generator) error {
		files, err := GetFilesInDir(fmt.Sprintf("%s/%s", CONTENT_DIR, POSTS_DIR))
		if err != nil {
			return err
		}

		var i int
		archive := make([]Archive, len(files))
		for _, file := range files {
			filename, filebyte, err := GetFileData(file)
			if err != nil {
				return err
			}

			markdown := goldmark.New(
				goldmark.WithExtensions(
					meta.Meta,
				),
			)

			var buf bytes.Buffer
			ctx := parser.NewContext()
			if err := markdown.Convert(filebyte, &buf, parser.WithContext(ctx)); err != nil {
				return err
			}

			metadata := meta.Get(ctx)
			id, err := strconv.Atoi(
				strings.TrimSuffix(
					strings.SplitAfter(metadata["published"].(string), "-")[0], "-",
				),
			)
			if err != nil {
				return err
			}

			if i == 0 {
				archive[i].ID = id
				archive[i].Posts = append(archive[i].Posts, Post{
					Title:       metadata["raw-title"].(string),
					Description: metadata["description"].(string),
					Author:      metadata["author"].(string),
					Published:   metadata["published"].(string),
					Keywords:    metadata["keywords"].(string),
					URL: template.URL(
						fmt.Sprintf(
							"%s/%s/%s.html",
							g.BaseURL,
							"posts",
							strings.TrimSuffix(filename, ".md"),
						),
					),
				})

				i++
			} else {
				if archive[i-1].ID != id {
					archive[i].ID = id
					archive[i].Posts = append(archive[i].Posts, Post{
						Title:       metadata["raw-title"].(string),
						Description: metadata["description"].(string),
						Author:      metadata["author"].(string),
						Published:   metadata["published"].(string),
						Keywords:    metadata["keywords"].(string),
						URL: template.URL(
							fmt.Sprintf(
								"%s/%s/%s.html",
								g.BaseURL,
								"posts",
								strings.TrimSuffix(filename, ".md"),
							),
						),
					})

					i++
				} else {
					archive[i-1].Posts = append(archive[i-1].Posts, Post{
						Title:       metadata["raw-title"].(string),
						Description: metadata["description"].(string),
						Author:      metadata["author"].(string),
						Published:   metadata["published"].(string),
						Keywords:    metadata["keywords"].(string),
						URL: template.URL(
							fmt.Sprintf(
								"%s/%s/%s.html",
								g.BaseURL,
								"posts",
								strings.TrimSuffix(filename, ".md"),
							),
						),
					})
				}
			}
		}

		sort.SliceStable(archive, func(i, j int) bool {
			return archive[j].ID < archive[i].ID
		})

		if len(archive) > 1 {
			g.Archive = archive[:len(archive)-1]
			return nil
		}

		g.Archive = archive
		return nil
	}
}
