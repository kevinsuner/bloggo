package render

import (
	"errors"

	"github.com/spf13/viper"
)

var RendererFunctions = map[string]RendererOption{
	"base-url":          WithBaseURL(),
	"theme":             WithTheme(),
	"index.posts.limit": WithPostsLimit(),
}

type RendererOption func(*Renderer) error

type Renderer struct {
	BaseURL    string
	Theme      string
	PostsLimit int
}

func New(opts ...RendererOption) *Renderer {
	r := new(Renderer)
	for _, opt := range opts {
		opt(r)
	}

	return r
}

func (r *Renderer) Render() error {
	return nil
}

func WithBaseURL() RendererOption {
	return func(r *Renderer) error {
		if len(viper.GetString("base-url")) == 0 {
			return errors.New("no base-url provided")
		}

		r.BaseURL = viper.GetString("base-url")
		return nil
	}
}

func WithTheme() RendererOption {
	return func(r *Renderer) error {
		if len(viper.GetString("theme")) == 0 {
			return errors.New("no theme provided")
		}

		r.Theme = viper.GetString("theme")
		return nil
	}
}

func WithPostsLimit() RendererOption {
	return func(r *Renderer) error {
		if viper.GetInt("index.posts.limit") <= 0 {
			return errors.New("no posts limit provided")
		}

		r.PostsLimit = viper.GetInt("index.posts.limit")
		return nil
	}
}
