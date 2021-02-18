package pages

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/yuin/goldmark"
)

// Page represents a page
type Page struct {
	Title, Author string
	Content       template.HTML
}

// NewPage creates a new Page
func NewPage(title, author, content string) *Page {
	md := goldmark.New()
	strB := &strings.Builder{}
	err := md.Convert([]byte(content), strB)
	if err != nil {
		panic(err)
	}

	return &Page{
		Title:   title,
		Author:  author,
		Content: template.HTML(strB.String()),
	}
}

// ID generates and ID for a Page
func (p *Page) ID() string {
	r := strings.NewReplacer(" ", "-", "?", " ")

	red := r.Replace(p.Title)

	return strings.ToLower(red)
}

// String creates a new templated page
func (p *Page) String() string {
	strBuilder := &strings.Builder{}
	tPage.Execute(strBuilder, p)

	fmt.Println("Rendered " + p.ID())

	return strBuilder.String()
}
