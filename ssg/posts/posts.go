package posts

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/manen/website/ssg/pages"
)

func ResolvePosts() []*pages.Page {
	dirStr := "./posts"
	dir, err := ioutil.ReadDir(dirStr)
	if err != nil {
		panic(err)
	}

	var posts []*pages.Page

	for _, v := range dir {
		if !strings.HasSuffix(v.Name(), ".md") {
			fmt.Println("Non-md file caught: " + v.Name())
			continue
		}

		pt := path.Join(dirStr, v.Name())
		posts = append(posts, ResolvePost(pt))
	}

	return posts
}

func ResolvePost(path string) *pages.Page {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return pages.NewPage("N/A", "N/A", string(bytes))
}
