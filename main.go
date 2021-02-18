package main

import (
	"github.com/manen/website/ssg/build"
	"github.com/manen/website/ssg/pages"
)

func main() {
	page := pages.NewPage("test page", "manen", `# test page but rendered

> Very WIP btw

## Render yes`)

	otherPageAndYesIRefuseToUseAnArray := pages.NewPage("test page but no", "joe", "joe who?")

	build.Build([]*pages.Page{page, otherPageAndYesIRefuseToUseAnArray})
}
