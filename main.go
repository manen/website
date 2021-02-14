package main

import (
	"fmt"

	"github.com/manen/website/ssg/pages"
)

func main() {
	page := pages.NewPage("test page", "manen", "hey, it actually worked?")

	fmt.Println(page.ToString())
}
