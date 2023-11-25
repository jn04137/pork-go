package controllers

import (
	"net/http"
	"log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"

	"banhbao/porkgo/db"
)

func AddPointToPostHandler(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "postId")
	data := &PointReq{}

	if bindErr := render.Bind(r, data); bindErr != nil {
		log.Printf("Binding error")
        render.Render(w, r, ErrNotFound)
        return
	}

	cookie, cookieErr := r.Cookie("jwtCookie")
	if cookieErr != nil {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, ErrNotFound)
        return
	}

	// have to update this with a secret for the jwt password
	_, err := jwt.ParseWithClaims(cookie.Value, &KnoAuthCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("testhash"), nil
	}); if err != nil {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, ErrNotFound)
        return
	}

	query := `INSERT INTO "PointsOnPost" VALUES ($1, $2, $3)"`

	_, qErr := db.DB.Exec(query, cookie.Value, postId, data.PointType); if qErr != nil {
		log.Printf("This is the query error: %v", qErr)
        render.Render(w, r, ErrNotFound)
        return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, successfulResponse)
}

func RemovePointFromPostHandler(w http.ResponseWriter, r *http.Request) {

}

func AddPointFromCommentHandler(w http.ResponseWriter, r *http.Request) {

}

func RemovePointFromCommentHandler(w http.ResponseWriter, r *http.Request) {

}

type Point struct {
	PointType string `json:"pointType"`
}

type PointReq struct {
	*Point
}

func (a *PointReq) Bind(r *http.Request) error {
	return nil
}
