package main

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/adrg/frontmatter"
)

type metadata struct {
	Lang          string `yaml:"lang"`
	Title         string `yaml:"title"`
	RawTitle      string `yaml:"raw_title"`
	Description   string `yaml:"description"`
	Keywords      string `yaml:"keywords"`
	Author        string `yaml:"author"`
	Robots        string `yaml:"robots"`
	OGType        string `yaml:"og_type"`
	Section       string `yaml:"section"`
	PublishedTime string `yaml:"published_time"`
	ModifiedTime  string `yaml:"modified_time"`
}

func parseMetadata(content []byte) ([]byte, *metadata, error) {
	var metadata *metadata
	rest, err := frontmatter.Parse(strings.NewReader(string(content)), &metadata)
	if err != nil {
		return nil, nil, err
	}

	return rest, metadata, nil
}

type templateData struct {
	Header  header
	Content template.HTML
	Posts   []post
	Archive map[int][]post
}

type header struct {
	Meta         *metadata
	BaseURL      template.URL
	CanonicalURL template.URL
	OGURL        template.URL
}

type post struct {
	ID            int
	Title         string
	Description   string
	Keywords      string
	Author        string
	PublishedTime string
	URL           template.URL
}

func parseTemplate(name string, files []string, data *templateData) ([]byte, error) {
	t, err := template.New(name).ParseFiles(files...)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
