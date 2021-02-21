package posts

import (
	"io/ioutil"
	"path"

	"github.com/manen/website/ssg/pages"
)

// ResolvePosts does all the required calculations to
// resolve all the data about all the posts
func ResolvePosts() []*pages.Page {
	var posts []*pages.Page
	for _, post := range data.Posts {
		bytes, err := ioutil.ReadFile(path.Join("./posts/content/", post.ID) + ".md")
		if err != nil {
			panic(err)
		}
		posts = append(posts, pages.NewPage(post.Title, post.ID, post.Author, string(bytes)))
	}

	return posts
}
