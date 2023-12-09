package controllers

import (
	"net/http"
	"log"
	"encoding/json"
	"strconv"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"

	"banhbao/porkgo/db"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	cookie, cookieErr := r.Cookie("jwtCookie")
	if cookieErr != nil {
		log.Printf("This is the cookieErr: %v", cookieErr)
		log.Printf("JWT Access Token was not found, this is the jwtCookie: %v", cookie)
		render.Render(w, r, ErrNotFound)
		return
	}

	data := &PostReq{}
	if bindErr := render.Bind(r, data); bindErr != nil {
		log.Printf("Binding error")
    render.Render(w, r, ErrNotFound)
    return
	}

	// Need token to create the post with correct user as owner
	token, err := jwt.ParseWithClaims(cookie.Value, &KnoAuthCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		log.Printf("This is the error: %v", err)
	}
	claims, ok := token.Claims.(*KnoAuthCustomClaims)

	if !token.Valid || !ok {
		render.Render(w, r, ErrNotFound)
		return
	}

	query := `INSERT INTO 
        "UserPost" (Owner, Title, Body) 
        VALUES ((SELECT ID from "UserAccount" WHERE Uuid=$1), $2, $3)`
	_, qErr := db.DB.Exec(query, claims.Uuid, data.Title, data.Body); if qErr != nil {
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
  var listOfPosts []PostModel = make([]PostModel, 0)

	query := `SELECT
				Username, Title, Body, "UserPost".CreatedAt, "UserPost".ID
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
	var postScans PostModel
		if err := rows.Scan(
			&postScans.Owner,
			&postScans.Title,
			&postScans.Body,
			&postScans.CreatedAt,
			&postScans.PostId); err != nil {
			log.Printf("Rows scan error: %v", err)
		}
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
		Username, Title, Body, "UserPost".CreatedAt, "UserPost".ID
		FROM "UserPost"
		INNER JOIN "UserAccount" ON "UserPost".Owner="UserAccount".ID
		WHERE "UserPost".ID=$1`

	var postScan PostModel
	row := db.DB.QueryRow(query, post)
	scanError := row.Scan(&postScan.Owner,
			&postScan.Title,
			&postScan.Body,
			&postScan.CreatedAt,
			&postScan.PostId)
	if scanError != nil {
		log.Printf("This is the scanError: %v", scanError)
	}

	responseData := map[string]interface{} {
		"message": "Everything went okay",
		"post": postScan,
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, responseData)
}

type PostModel struct {
	Title	    string
	Owner	    string
	Body	    string
	PostId		string
  CreatedAt string
  Count     int
}

type Post struct {
	Title	string 			`json:"title"`
	Body	json.RawMessage `json:"body"`
}

type PostReq struct {
	*Post
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

