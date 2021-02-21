package posts

import (
	"github.com/manen/website/ssg/pages"
)

func ResolvePosts() []*pages.Page {
	var posts []*pages.Page
	for _, post := range data.Posts {
		posts = append(posts, pages.NewPage(post.Title, post.Author, "uh well now this is unimplemented so"))
	}

	return posts
}
