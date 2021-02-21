package build

import (
	"fmt"
	"io/ioutil"

	"github.com/manen/website/ssg/pages"
)

// Build is the main build function
func Build(ps []*pages.Page) {
	copyPublic()
	for _, p := range ps {
		err := ioutil.WriteFile("./build/"+p.ID+".html", []byte(p.String()), 0644) // Should I know what 0644 is? probably
		if err != nil {
			fmt.Println(err)
		}
	}

	index := pages.NewIndex(ps)

	err := ioutil.WriteFile("./build/"+index.ID+".html", []byte(index.String()), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
