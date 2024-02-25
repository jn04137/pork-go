package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"porkgo/db"
	"porkgo/porkmiddleware"
	"porkgo/models"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	jwtUserInfo := r.Context().Value("jwtUserValues").(porkmiddleware.JwtUserValues)
	data := &models.PostReq{}
	if bindErr := render.Bind(r, data); bindErr != nil {
		log.Printf("Binding error")
		render.Render(w, r, ErrNotFound)
		return
	}

	query := `INSERT INTO
        "UserPost" (Owner, Title, Body) 
        VALUES ((SELECT ID from "UserAccount" WHERE Uuid=$1), $2, $3)`
	_, qErr := db.DB.Exec(query, jwtUserInfo.Uuid, data.Title, data.Body); if qErr != nil {
		log.Printf("This is the query error: %v", qErr)
        render.Render(w, r, ErrNotFound)
        return
	}

	responseData := map[string]interface{} {
		"message": "was a success",
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, responseData)
}

// Should have the main feed be tailored for each end user
func LoadFeedPost(w http.ResponseWriter, r *http.Request) {
  page, parseErr := strconv.Atoi(chi.URLParam(r, "page")); if parseErr != nil {
    log.Printf("parse error")
  }
  var listOfPosts []models.PostModel = make([]models.PostModel, 0)

	query := `SELECT 
			Username, Title, Body, "UserPost".CreatedAt, "UserPost".ID, "UserPost".isDeleted
			FROM "UserPost"
			INNER JOIN "UserAccount"
			ON "UserAccount".ID="UserPost".Owner
			ORDER BY "UserPost".CreatedAt DESC
			LIMIT 10
			OFFSET $1`

	rows, qErr := db.DB.Query(query, page); if qErr != nil {
		log.Printf("This is the query error: %v", qErr)
		render.Render(w, r, ErrNotFound)
		return
	}

	for rows.Next() {
		var postScans models.PostModel
			err := rows.Scan(
				&postScans.Owner,
				&postScans.Title,
				&postScans.Body,
				&postScans.CreatedAt,
				&postScans.PostId,
				&postScans.IsDeleted); 
			if err != nil {
				log.Printf("Rows scan error: %v", err)
				continue
			}
			isPostDeleted(&postScans)
		listOfPosts = append(listOfPosts, postScans)
	}
  rows.Close()
  
  nextCursor := page + 10
  responseData := map[string]interface{} {
    "posts": listOfPosts,
    "nextCursor": nextCursor,
  }

  if len(listOfPosts) == 0 {
    responseData["nextCursor"]=nil
  }

	render.Status(r, http.StatusOK)
	render.JSON(w, r, responseData)
}

func ViewPost(w http.ResponseWriter, r *http.Request) {
	post := chi.URLParam(r, "postId")

	query := `SELECT
		Username, Title, Body, "UserPost".CreatedAt, "UserPost".ID, "UserPost".isDeleted
		FROM "UserPost"
		INNER JOIN "UserAccount" ON "UserPost".Owner="UserAccount".ID
		WHERE "UserPost".ID=$1`

	var postScan models.PostModel
	row := db.DB.QueryRow(query, post)
	scanError := row.Scan(&postScan.Owner,
			&postScan.Title,
			&postScan.Body,
			&postScan.CreatedAt,
			&postScan.PostId,
			&postScan.IsDeleted)
	if scanError != nil {
		log.Printf("This is the scanError: %v", scanError)
	}
	isPostDeleted(&postScan)

	responseData := map[string]interface{} {
		"message": "Everything went okay",
		"post": postScan,
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, responseData)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {

	jwtUserInfo := r.Context().Value("jwtUserValues").(porkmiddleware.JwtUserValues)
	var deletePost models.PostModel
	err := json.NewDecoder(r.Body).Decode(&deletePost)
	if err != nil {
		log.Printf("This is the error: %v", err)
		render.Render(w, r, ErrNotFound)
		return
	}

	query := `UPDATE "UserPost" SET isDeleted=TRUE WHERE ID=$1
		AND Owner=(SELECT ID FROM "UserAccount" WHERE Uuid=$2)`

	_, qErr := db.DB.Exec(query, deletePost.PostId, jwtUserInfo.Uuid); if qErr != nil {
		log.Printf("This is the query error: %v", qErr)
		render.Render(w, r, ErrNotFound)
		return
	}

	responseData := map[string]string {
		"message": "post was deleted",
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, responseData)

}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	jwtUserInfo := r.Context().Value("jwtUserValues").(porkmiddleware.JwtUserValues)
	var editPost models.PostModel
	err := json.NewDecoder(r.Body).Decode(&editPost)
	log.Printf("This is the json data: %v", editPost)
	if err != nil {
		log.Printf("This is the error: %v", err)
		render.Render(w, r, ErrNotFound)
		return
	}
	query := `UPDATE "UserPost" SET Body=$1 WHERE ID=$2
		AND Owner=(SELECT ID FROM "UserAccount" WHERE Uuid=$3)`
	
	_, qErr := db.DB.Exec(query, editPost.Body, editPost.PostId, jwtUserInfo.Uuid); if qErr != nil {
		log.Printf("This is the query error: %v", qErr)
		render.Render(w, r, ErrNotFound)
		return
	}

	responseData := map[string]string {
		"message": "post was changed",
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, responseData)
}

// Checks if post has deleted flag
// post prints as if deleted if flag is true
func isPostDeleted(post *models.PostModel) *models.PostModel {
	if post.IsDeleted {
		post.Title = "Deleted Post"
		post.Body = "Deleted Post"
		post.Owner = "Anonymous User"
	}

	return post
}

