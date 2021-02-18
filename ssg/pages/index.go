package pages

import (
	"fmt"
	"strings"
)

// Index represents the landing page
type Index struct {
	Pages []*Page
}

// NewIndex creates a new Index
func NewIndex(pages []*Page) *Index {
	return &Index{
		Pages: pages,
	}
}

// ID returns "index"
func (i *Index) ID() string {
	return "index"
}

// String generates the page
func (i *Index) String() string {
	strBuilder := &strings.Builder{}
	tIndex.Execute(strBuilder, i)

	fmt.Println("Rendered " + i.ID())

	return strBuilder.String()
}
