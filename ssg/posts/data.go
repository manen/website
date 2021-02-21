package posts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Data struct {
	Posts []*PostData `json:"posts"`
}

type PostData struct {
	Title  string `json:"title"`
	ID     string `json:"id"`
	Author string `json:"author"`
}

var data Data

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
