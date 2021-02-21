package posts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Data represents data about the website
// (./posts/data.json)
type Data struct {
	Posts []*PostData `json:"posts"`
}

// PostData is data about a post
type PostData struct {
	Title  string `json:"title"`
	ID     string `json:"id"`
	Author string `json:"author"`
}

// data is an initialized copy of ./posts/data.json
// (initialized in init)
var data Data

// init initialized data
func init() {
	bytes, err := ioutil.ReadFile("./posts/data.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
