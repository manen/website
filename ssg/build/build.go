package build

import (
	"fmt"
	"io/ioutil"

	"github.com/manen/website/ssg/pages"
)

// Build is the main build function
func Build(pages []*pages.Page) {
	copyPublic()
	for _, p := range pages {
		err := ioutil.WriteFile("./build/"+p.ID()+".html", []byte(p.String()), 0644) // Should I know what 0644 is? probably
		if err != nil {
			fmt.Println(err)
		}
	}
}
