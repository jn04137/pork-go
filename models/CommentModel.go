package models

import (
	"net/http"
)

type CommentModel struct {
	Owner     string
	Body      string
	PostId    int
	CreatedAt string
}

type Comment struct {
	Body string `json:"body"`
}

type CommentReq struct {
	*Comment
}

func (a *CommentReq) Bind(r *http.Request) error {
	return nil
}
