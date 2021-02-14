package pages

import (
	"html/template"
	"io/ioutil"
)

var (
	tPage *template.Template
)

func init() {
	page, err := ioutil.ReadFile("./pages/page.html")
	if err != nil {
		panic(err)
	}

	pageStr := string(page)

	tPage, err = template.New("page").Parse(pageStr)
	if err != nil {
		panic(err)
	}
}
