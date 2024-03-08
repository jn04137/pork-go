package models

import (
	"net/http"
)

type PostModel struct {
	Title	    	string
	Owner	    	string
	Body	    	string
	PostId			string
	CreatedAt 	string
	Count     	int
	IsDeleted		bool
}

type PostReq struct {
	*PostModel
}

func (a *PostReq) Bind(r *http.Request) error {
	return nil
}

type ListOfPostRes struct {
  ListOfPosts []PostModel `json:"posts"` 
}

func (a *ListOfPostRes) Bind(r *http.Request) error {
	return nil
}
