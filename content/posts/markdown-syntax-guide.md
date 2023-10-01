---
lang: "en"
title: "Markdown Syntax Guide | Bloggo"
raw_title: "Markdown Syntax Guide"
description: "This article offers a sample of basic Markdown syntax that can be used in Bloggo content files, also it shows whether basic HTML elements are decorated with CSS in a Bloggo theme."
keywords: "bloggo, go, golang, ssg, markdown, frontmatter, simple, minimalist"
author: "itsksrof"
robots: "index, follow"
og_type: "article"
section: "posts"
published_time: "2023-09-27"
modified_time: "2023-09-27"
---
# Markdown Syntax Guide
This article offers a sample of basic Markdown syntax that can be used in Bloggo content files, also it shows whether basic HTML elements are decorated with CSS in a Bloggo theme.

## Headings
The following HTML `<h1>`-`<h6>` elements represent six levels of section headings. `<h1>` is the highest section level while `<h6>` is the lowest.

# H1
## H2
### H3
#### H4
##### H5
###### H6

## Paragraph
Suspendisse potenti. Aenean placerat enim nulla, quis volutpat neque hendrerit nec. Nam laoreet eu ante vel placerat. Sed non orci ornare, euismod diam ut, convallis felis. Maecenas molestie, justo eget faucibus fringilla, odio lectus luctus lectus, sed semper eros erat nec ipsum. Fusce a arcu nulla. Aliquam suscipit erat sit amet suscipit suscipit. Cras sollicitudin arcu eu diam vestibulum, sit amet molestie purus sollicitudin. Nulla lobortis nulla eu orci imperdiet accumsan. Integer porta, diam eget elementum ullamcorper, ligula purus egestas lorem, convallis faucibus ligula neque id justo.

Pellentesque quis rutrum velit. Etiam odio ipsum, scelerisque a nunc eu, blandit laoreet turpis. Nam molestie tincidunt arcu, eu pretium metus egestas vehicula. In id consectetur dui. Etiam dictum, mauris eu lobortis imperdiet, sapien ipsum fringilla ipsum, sit amet lobortis tortor nisl vitae dolor. Quisque porttitor tristique velit.

## Image
You can use the following syntax to include an image. Path of the image should be relative to the `public` folder.

```markdown
![Matterhorn](/bloggo/assets/1.jpg)
```

![Matterhorn](/bloggo/assets/1.jpg)

You can also include images from external sources.

```markdown
![Ferrari](https://source.unsplash.com/random/640x410/?ferrari)
```

![Ferrari](https://source.unsplash.com/random/640x410/?ferrari)

## Blockquotes
The blockquote element represents content that is quoted from another source, optionally with a citation which must be within a `footer` or `cite` element, and optionally with in-line changes such as annotations and abbreviations.

### Blockquote without attribution
> You can use Markdown syntax within a blockquote, like **bold**, *italics*, [links](https://ksrof.com), `code`.

### Blockquote with attribution
> Talk is cheap. Show me the code.
>
> *Linus Torvalds*

## Code Blocks

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

### Inline code
Use the backticks to refer to a `variable` within a sentence.

## List Types

### Ordered List
1. First item
2. Second item
3. Third item

### Unordered List
- List item
- Another item
- And another item

### Nested list
- Fruit
    - Apple
    - Orange
    - Banana
- Dairy
    - Milk
    - Cheese

I've used [Markdown Syntax Guide](https://hugo-blog-awesome.netlify.app/posts/markdown-syntax/) from [Hugo Awesome Blog](https://github.com/hugo-sid/hugo-blog-awesome) as a benchmark to see how many features does Bloggo support.
The beautiful image of the Matterhorn was made by [Chris Holgersson](https://unsplash.com/@chrisholgersson) from Unsplash.