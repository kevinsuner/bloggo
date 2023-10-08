<h1 align="center">Bloggo</h1>
<h4 align="center">🚀 Fast | 🪶 Lightweight | 📱 Responsive</h4>
<br/>

Bloggo is a free and simplistic [static site generator](https://en.wikipedia.org/wiki/Static_site_generator) for individuals that prioritize content over features, it lets you use custom themes by running a simple command, and uses [Go](https://go.dev), Markdown and Frontmatter to generate HTML pages that are optimized for performance and SEO.

[![Go Version](https://img.shields.io/github/go-mod/go-version/itsksrof/bloggo)](https://github.com/itsksrof/bloggo/blob/master/go.mod)
[![Pagespeed Insights](https://img.shields.io/badge/pagespeed-insights-green)](https://pagespeed.web.dev/report?url=https://itsksrof.github.io/bloggo/)
[![License](https://img.shields.io/github/license/itsksrof/bloggo)](https://github.com/itsksrof/bloggo/blob/master/LICENSE)

---

<p align="center">
    <kbd><img src="https://i.postimg.cc/wTjWHF0j/bloggo.jpg" alt="Bloggo Theme" title="Bloggo Theme"/></kbd>
</p>

---

## Usage
The following guide will start a local Fiber server at `http://127.0.0.1:5500/public`, will use the default theme provided with Bloggo, and will limit the amount of articles shown in the homepage to 20.

1. Clone this repository
    ```bash
    git clone https://github.com/itsksrof/bloggo.git
    ```
2. Go to the cloned repository directory
    ```bash
    cd /path/to/bloggo
    ```
3. Go to the `bloggo.yaml` file and set `base-url` and `test` values to the following
    ```yaml
    base-url: http://127.0.0.1:5500/public
    test: true
    ```
4. Run the following command and head to `http://127.0.0.1:5500/public` on your browser
    ```bash
    # For Linux users
    ./bin/bloggo-amd64-linux
    # For Windows users
    ./bin/bloggo-amd64.exe
    # For MacOS users
    ./bin/bloggo-amd64-darwin
    ```

### Configuration
The following configuration options can be set in the `bloggo.yaml` file.

| Name | Required | Example | Description |
| ---- | -------- | ----- | ----------- |
| lang | true | `lang: "en"` | The default language of the site |
| base-url | true | `base-url: "http://127.0.0.1:5500/public"` | The base URL used in meta-tags and links |
| theme | true | `theme: "bloggo"` | The theme used across the site |
| posts-limit | true |  `posts-limit: 10` | The limit of articles shown in the index page |
| test | true | `test: false` | Preview of the site using a server with Fiber |
| pages | true | `pages: ["index", "about", "archive", "404"]` | The pages that are going to be generated |
| dirs | true | `dirs: ["assets", "css"]` | The directories that are going to be copied from the `theme` folder to the `public` folder |

### Frontmatter
The following meta-tags can be used at the top of each `.md` file like this:
```yaml
---
title: "Markdown Syntax Guide | Bloggo"
raw-title: "Markdown Syntax Guide"
description: "This article offers a sample of basic Markdown syntax that can be used in Bloggo content files, also it shows whether basic HTML elements are decorated with CSS in a Bloggo theme."
keywords: "bloggo, go, golang, ssg, markdown, frontmatter, simple, minimalist"
author: "itsksrof"
robots: "index, follow"
type: "article"
section: "posts"
published: "2023-09-27"
modified: "2023-09-27"
---
```

| Name | Required | Usage | Description |
| ---- | -------- | ----- | ----------- |
| lang | True | `lang: "es"` | The language of the page |
| title | True | `title: "Home • Bloggo"` | The title of the page |
| raw-title | Only in posts | `raw-title: "Home"` | The raw title of the page |
| description | True | `description: "The description"` | The description of the page |
| keywords | True | `keywords: "foo, bar, foobar"` | The keywords of the page |
| author | True | `author: "itsksrof"` | The author of the page |
| robots | True | `robots: "index, follow"` | The instructions for the search engine |
| type | True | `type: "website"` | The type of the object for Open Graph |
| section | Only in posts | `section: "posts"` | The section in which the article belongs |
| published | Only in posts | `published: "2023-09-30"` | The date in which the article was published |
| modified | Only in posts | `modified: "2023-09-30"` | The date in which the article was modified |

## Theming
Given that Bloggo focuses on content over features, theming is very straightforward, as there is no possibility of generating more pages, unless you want to dig in the source code to do so. This is by design and I wish to keep it that way. Therefore to add a new theme, you only need to respect a certain file structure, this doesn't mean that you can't add new things, for example you've could create a navigation bar that included links to your social media, or a homepage with an avatar if you wish, so yes, you are certainly limited, but still you can do a lot of things within those bounds. If Bloggo doesn't suit your use-case there are wonderful alternatives like [Hugo](https://gohugo.io/) that will surely do.

As a recommendation, you can use the bloggo theme file structure as a reference.
```text
bloggo/
    assets/
    css/
    partials/
        footer.html
        header.html
        nav.html
    about.html
    archive.html
    index.html
    post.html
```

## Deploying
At the moment the only proven way to deploy Bloggo is through [Github Pages](https://pages.github.com/). This will surely change in the future as I try to deploy Bloggo in other services such as Netlify, Linode, etc...

### Deploy with Github Pages
Bloggo uses a custom [GitHub Action](https://docs.github.com/en/actions/learn-github-actions/understanding-github-actions) located in [`.github/workflows`](https://github.com/itsksrof/bloggo/blob/master/.github/workflows/static.yml) to deploy to GitHub Pages, but, before you can start actually using it, you must tell your repo specifically to build your site that way. To do so you just have to:
1. On Github, navigate to your site's repository
2. Under your repository name, click **Settings**. If you cannot see the "Settings" tab, select the **•••** dropdown menu, then click **Settings**
3. In the "Code and Automation" section of the sidebar, click **Pages**
4. Under "Build and Deployment", under "Source", select **Github Actions**

In case you want to build your own action, you're having some issues, or you just want to know more head to [Configuring a publishing source for your GitHub Pages site](https://docs.github.com/en/pages/getting-started-with-github-pages/configuring-a-publishing-source-for-your-github-pages-site) or [Using custom workflows with GitHub Pages](https://docs.github.com/en/pages/getting-started-with-github-pages/using-custom-workflows-with-github-pages).

## Contributing
All contributions are extremely appreciated, even though the capabilities of Bloggo are intended to be just a few, that doesn't mean that there is no room for improvement, because there is, so if you find places where things can be improved don't hesitate and open a pull request or an issue, as long as things are under Bloggo's scope I will be delighted to work on them, or allow other contributors to do so.

Aside from that, there are a couple of things that need to be respected in order to have something merged into Bloggo.
- Improvements over features
- Fit naturally into Bloggo's scope

### Git commit message guidelines
The most important part is that each commit messages should have a title/subject in imperative mood starting with a capital letter and no trailing period: `generator: Fix articles list being sorted in ascending order` **NOT** `list sorted right.`. If you still unsure about how to write a good commit message this [blog article](https://cbea.ms/git-commit/) is a good resource for learning that.

Most title/subjects should have a lower-cased prefix with a colon and one whitespace. The prefix can be:
- The name of the package where (most of) the changes are made (e.g. `parse: Add RawTitle to metadata struct`)
- If the commit touches several packages with a common functional topic, use that as a prefix (e.g. `errors: Resolve correct line numbers`)
- If the commit touches several packages without a common functional topic, prefix with `all:` (e.g. `all: Reformat go code`)
- If this is a documentation update, prefix with (e.g. `docs:`)
- If nothing of the above applies, just leave the prefix out

Also, if your commit references some or more Github issues, always end your commit message body with *See #1234* or *Fixes #1234*.
Replace *1234* with the Github issue ID.

An example:
```text
generator: Fix articles list being sorted in ascending order

Added a function that returns a sorted "map" by passing it an unsorted one, 
it creates an "array" with the keys of the first "map", sorts the given "array"
using sort.StableSort and iterates over those keys to create new entries in the
new "map" and sets its value by accessing the first "map".

Fixes #1234
```

### Fetching the source from Github
Bloggo uses the Go Modules support built into Go 1.11 to build. The easiest is to clone Hugo in a directory outside of `GOPATH`, as in the following example:
```bash
mkdir $HOME/src
cd $HOME/src
git clone https://github.com/itsksrof/bloggo.git
cd bloggo
```

Now, to make a change to Bloggo's source:
1. Create a new branch for your changes (the branch name is arbitrary):
    ```bash
    git checkout -b abc123
    ```
2. After making your changes, commit them to your new branch:
    ```bash
    git commit -a -v
    ```
3. Fork Bloggo in Github
4. Add your fork as a new remote (the remote name, "foo" in this example, is arbitrary):
    ```bash
    git remote add foo git@github.com:USERNAME/bloggo.git
    ```
5. Push the changes to your new remote:
    ```bash
    git push --set-upstream foo abc123
    ```
6. You are now ready to submit a PR based upon the new branch in your forked repository

### Building Bloggo with your changes
Bloggo doesn't make use of any tools to build it (at the moment). You must run the following commands from the Bloggo working directory.
```bash
cd $HOME/src/bloggo
```

To build Bloggo:
```bash
# Linux 64-bit
GOOS=linux GOARCH=amd64 go build -o bin/bloggo-amd64-linux .
# Windows 64-bit
GOOS=windows GOARCH=amd64 go build -o bin/bloggo-amd64.exe .
# MacOS 64-bit
GOOS=darwin GOARCH=amd64 go build -o bin/bloggo-amd64-darwin .
```

## Credits

Bloggo makes use or has taken inspiration from a variety of open source projects including:
- [github.com/yuin/goldmark](github.com/yuin/goldmark)
- [github.com/yuin/goldmark-meta](github.com/yuin/goldmark-meta)
- [github.com/gofiber/fiber/v2](github.com/gofiber/fiber/v2)
- [github.com/gohugoio/hugo](github.com/gohugoio/hugo)
- [github.com/egonelbre/gophers](github.com/egonelbre/gophers)