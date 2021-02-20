package main

import (
	"github.com/manen/website/ssg/build"
	"github.com/manen/website/ssg/posts"
)

func main() {
	build.Build(posts.ResolvePosts())
}
