package lib

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	Id        string `json:"id"`
	Author    string `json:"author"`
	Timestamp int64  `json:"timestamp"`
}

type Postfile struct {
	Filetype string `json:"filetype"`
	Version  string `json:"version"`
	Posts    []Post `json:"posts"`
}

func FormatPost(post Post) Post {
	if post.Timestamp == 0 {
		post.Timestamp = time.Now().Unix()
	}

	if post.Id == "" {
		post.Id = uuid.New().String()
	}

	return post
}
