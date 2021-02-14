package pages

import "strings"

// Page represents a page
type Page struct {
	Title, Author, Content string
}

// NewPage creates a new Page
func NewPage(Title, Author, Content string) *Page {
	return &Page{
		Title, Author, Content,
	}
}

// ToString creates a new templated page
func (p *Page) ToString() string {
	strBuilder := &strings.Builder{}
	tPage.Execute(strBuilder, p)

	return strBuilder.String()
}
