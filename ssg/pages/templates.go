package pages

import (
	"html/template"
	"io/ioutil"
)

var (
	tPage  *template.Template
	tIndex *template.Template
)

func init() {
	page, err := ioutil.ReadFile("./pages/page.html")
	if err != nil {
		panic(err)
	}
	index, err := ioutil.ReadFile("./pages/index.html")
	if err != nil {
		panic(err)
	}

	pageStr := string(page)
	indexStr := string(index)

	tPage, err = template.New("page").Parse(pageStr)
	if err != nil {
		panic(err)
	}
	tIndex, err = template.New("index").Parse(indexStr)
	if err != nil {
		panic(err)
	}
}
