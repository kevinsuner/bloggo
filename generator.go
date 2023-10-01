// SPDX-License-Identifier: MIT
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/yuin/goldmark"
)

type generator struct {
	baseURL      string
	themeDir     string
	canonicalURL string
	partialsDir  string
}

func newGenerator(baseURL, themeDir string) *generator {
	return &generator{
		baseURL:      baseURL,
		themeDir:     themeDir,
		canonicalURL: fmt.Sprintf("%s/%s", baseURL, "index.html"),
		partialsDir:  fmt.Sprintf("%s/%s/%s", THEMES_DIR, themeDir, "partials"),
	}
}

func (g *generator) index(limit int) error {
	files, err := getFilesInDir(CONTENT_POSTS_DIR)
	if err != nil {
		return err
	}

	posts := []post{}
	for i, file := range files {
		if i <= limit {
			filename, filebyte, err := getFileData(file)
			if err != nil {
				return err
			}

			_, metadata, err := parseMetadata(filebyte)
			if err != nil {
				return err
			}

			date, err := strconv.Atoi(
				strings.ReplaceAll(metadata.PublishedTime, "-", ""),
			)
			if err != nil {
				return err
			}

			posts = append(posts, post{
				ID:            date, // use date as an id for sorting
				Title:         metadata.RawTitle,
				Description:   metadata.Description,
				Keywords:      metadata.Keywords,
				Author:        metadata.Author,
				PublishedTime: metadata.PublishedTime,
				URL:           template.URL(g.setURL("posts", "", filename)),
			})
		}
	}

	sort.SliceStable(posts, func(i, j int) bool {
		return posts[j].ID < posts[i].ID
	})

	filename, filebyte, err := getFileData(fmt.Sprintf("%s/%s.md", CONTENT_DIR, "index"))
	if err != nil {
		return err
	}

	content, metadata, err := parseMetadata(filebyte)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(content, &buf); err != nil {
		return err
	}

	tmpl, err := parseTemplate(
		"index",
		[]string{
			fmt.Sprintf("%s/%s/%s.html", THEMES_DIR, g.themeDir, "index"),
			fmt.Sprintf("%s/%s.html", g.partialsDir, "header"),
			fmt.Sprintf("%s/%s.html", g.partialsDir, "nav"),
			fmt.Sprintf("%s/%s.html", g.partialsDir, "footer"),
		},
		&templateData{
			Header: header{
				Meta:         metadata,
				BaseURL:      template.URL(g.baseURL),
				CanonicalURL: template.URL(g.canonicalURL),
				OGURL:        template.URL(g.setURL("", "", filename)),
			},
			Content: template.HTML(buf.String()),
			Posts:   posts,
		},
	)
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("./%s/%s.html", PUBLIC_DIR, filename),
		tmpl,
		0644,
	); err != nil {
		return err
	}

	return nil
}

func (g *generator) page(name string) error {
	filename, filebyte, err := getFileData(fmt.Sprintf("%s/%s.md", CONTENT_DIR, name))
	if err != nil {
		return err
	}

	content, metadata, err := parseMetadata(filebyte)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(content, &buf); err != nil {
		return err
	}

	tmpl, err := parseTemplate(
		name,
		[]string{
			fmt.Sprintf("%s/%s/%s.html", THEMES_DIR, g.themeDir, name),
			fmt.Sprintf("%s/%s.html", g.partialsDir, "header"),
			fmt.Sprintf("%s/%s.html", g.partialsDir, "nav"),
			fmt.Sprintf("%s/%s.html", g.partialsDir, "footer"),
		},
		&templateData{
			Header: header{
				Meta:         metadata,
				BaseURL:      template.URL(g.baseURL),
				CanonicalURL: template.URL(g.canonicalURL),
				OGURL:        template.URL(g.setURL("", "", filename)),
			},
			Content: template.HTML(buf.String()),
		},
	)
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("./%s/%s.html", PUBLIC_DIR, filename),
		tmpl,
		0644,
	); err != nil {
		return err
	}

	return nil
}

func (g *generator) archive() error {
	files, err := getFilesInDir(CONTENT_POSTS_DIR)
	if err != nil {
		return err
	}

	posts := []post{}
	for _, file := range files {
		filename, filebyte, err := getFileData(file)
		if err != nil {
			return err
		}

		_, metadata, err := parseMetadata(filebyte)
		if err != nil {
			return err
		}

		date, err := strconv.Atoi(
			strings.ReplaceAll(metadata.PublishedTime, "-", ""),
		)
		if err != nil {
			return err
		}

		posts = append(posts, post{
			ID:            date, // use date as an id for sorting
			Title:         metadata.RawTitle,
			Description:   metadata.Description,
			Keywords:      metadata.Keywords,
			Author:        metadata.Author,
			PublishedTime: metadata.PublishedTime,
			URL:           template.URL(g.setURL("posts", "", filename)),
		})
	}

	sort.SliceStable(posts, func(i, j int) bool {
		return posts[j].ID < posts[i].ID
	})

	archive := make(map[int][]post)
	for _, p := range posts {
		year, err := strconv.Atoi(
			strings.TrimSuffix(
				strings.SplitAfter(p.PublishedTime, "-")[0], "-",
			),
		)
		if err != nil {
			return err
		}

		archive[year] = append(archive[year], post{
			ID:            p.ID,
			Title:         p.Title,
			Description:   p.Description,
			Keywords:      p.Keywords,
			Author:        p.Author,
			PublishedTime: p.PublishedTime,
			URL:           p.URL,
		})
	}

	filename, filebyte, err := getFileData(fmt.Sprintf("%s/%s.md", CONTENT_DIR, "archive"))
	if err != nil {
		return err
	}

	content, metadata, err := parseMetadata(filebyte)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(content, &buf); err != nil {
		return err
	}

	tmpl, err := parseTemplate(
		"archive",
		[]string{
			fmt.Sprintf("%s/%s/%s.html", THEMES_DIR, g.themeDir, "archive"),
			fmt.Sprintf("%s/%s.html", g.partialsDir, "header"),
			fmt.Sprintf("%s/%s.html", g.partialsDir, "nav"),
			fmt.Sprintf("%s/%s.html", g.partialsDir, "footer"),
		},
		&templateData{
			Header: header{
				Meta:         metadata,
				BaseURL:      template.URL(g.baseURL),
				CanonicalURL: template.URL(g.canonicalURL),
				OGURL:        template.URL(g.setURL("", "", filename)),
			},
			Content: template.HTML(buf.String()),
			Archive: archive,
		},
	)
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("./%s/%s.html", PUBLIC_DIR, filename),
		tmpl,
		0644,
	); err != nil {
		return err
	}

	return nil
}

func (g *generator) posts() error {
	files, err := getFilesInDir(CONTENT_POSTS_DIR)
	if err != nil {
		return err
	}

	for _, file := range files {
		filename, filebyte, err := getFileData(file)
		if err != nil {
			return err
		}

		content, metadata, err := parseMetadata(filebyte)
		if err != nil {
			return err
		}

		var buf bytes.Buffer
		if err := goldmark.Convert(content, &buf); err != nil {
			return err
		}

		tmpl, err := parseTemplate(
			"post",
			[]string{
				fmt.Sprintf("%s/%s/%s.html", THEMES_DIR, g.themeDir, "post"),
				fmt.Sprintf("%s/%s.html", g.partialsDir, "header"),
				fmt.Sprintf("%s/%s.html", g.partialsDir, "nav"),
				fmt.Sprintf("%s/%s.html", g.partialsDir, "footer"),
			},
			&templateData{
				Header: header{
					Meta:         metadata,
					BaseURL:      template.URL(g.baseURL),
					CanonicalURL: template.URL(g.canonicalURL),
					OGURL:        template.URL(g.setURL("posts", "", filename)),
				},
				Content: template.HTML(buf.String()),
			},
		)
		if err != nil {
			return err
		}

		if err := os.WriteFile(
			fmt.Sprintf("./%s/%s.html", PUBLIC_POSTS_DIR, filename),
			tmpl,
			0644,
		); err != nil {
			return err
		}
	}

	return nil
}

func (g *generator) setURL(folder, subfolder, filename string) string {
	if folder != "" {
		if subfolder != "" {
			return fmt.Sprintf(
				"%s/%s/%s/%s",
				g.baseURL,
				folder,
				subfolder,
				fmt.Sprintf("%s.html", filename),
			)
		}

		return fmt.Sprintf(
			"%s/%s/%s",
			g.baseURL,
			folder,
			fmt.Sprintf("%s.html", filename),
		)
	}

	return fmt.Sprintf(
		"%s/%s",
		g.baseURL,
		fmt.Sprintf("%s.html", filename),
	)
}
