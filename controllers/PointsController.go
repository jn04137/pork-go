package controllers

import (
	"net/http"
	"log"
  "context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"

	"banhbao/porkgo/db"
)

func GetPostPoints(w http.ResponseWriter, r *http.Request) {
  postId := chi.URLParam(r, "postId")
  //data := &PointReq{}
  var positiveCount int
  var negativeCount int

  positivePointsQ := `SELECT COUNT(*) FROM "PointsOnPost" WHERE PostId=$1 AND Point='plus'`
  negativePointsQ := `SELECT COUNT(*) FROM "PointsOnPost" WHERE PostId=$1 AND Point='minus'`


  ctx := context.Background()
  tx, txErr := db.DB.BeginTx(ctx, nil)
  if txErr != nil {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, ErrNotFound)
    return
  }
  defer tx.Rollback()
	
  qErr := tx.QueryRow(positivePointsQ, postId).Scan(&positiveCount)
  if qErr != nil {
		log.Printf("(line 32) This is the query error: %v", qErr)
    render.Render(w, r, ErrNotFound)
    return
	}
  
  qErr = tx.QueryRow(negativePointsQ, postId).Scan(&negativeCount)
  if qErr != nil {
		log.Printf("(line 38) This is the query error: %v", qErr)
    render.Render(w, r, ErrNotFound)
    return
	}

  if txErr = tx.Commit(); txErr != nil {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, ErrNotFound)
    return
  } 

  log.Printf("positive query: %v", positiveCount)
  log.Printf("negative query: %v", negativeCount)

  response := map[string]string{
    "points": "0",
  }

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}

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
