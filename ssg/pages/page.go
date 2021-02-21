package pages

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/yuin/goldmark"
)

// Page represents a page
type Page struct {
	Title, ID, Author string
	Content           template.HTML
}

// NewPage creates a new Page
func NewPage(title, id, author, content string) *Page {
	md := goldmark.New()
	strB := &strings.Builder{}
	err := md.Convert([]byte(content), strB)
	if err != nil {
		panic(err)
	}

	return &Page{
		Title:   title,
		ID:      id,
		Author:  author,
		Content: template.HTML(strB.String()),
	}
}

// String creates a new templated page
func (p *Page) String() string {
	strBuilder := &strings.Builder{}
	tPage.Execute(strBuilder, p)

	fmt.Println("Rendered " + p.ID)

	return strBuilder.String()
}
