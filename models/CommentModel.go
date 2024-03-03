package models

import (
	"net/http"
)

type CommentModel struct {
	Owner     	string
	Body      	string
	CommentId		string
	PostId    	int
	CreatedAt 	string
	IsDeleted		bool
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
