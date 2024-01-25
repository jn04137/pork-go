package controllers

import (
	"net/http"
  "fmt"
	"log"
  "context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

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

  points := fmt.Sprintf("%d", (positiveCount - negativeCount))
  response := map[string]string{
    "points": points,
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

  /*
    Check if the user has already liked the post. 
      - If they have liked it then change to 'empty' status
      - If they have disliked it, then like it
      - If it is empty or not existing, leave a new like
  */

  // initial searching for an existing record if it exists
  userPostPointQuery := `SELECT Point FROM "PointsOnPost" WHERE UserId=$1 AND PostId=$2`

  var userPointing string
  qErr :=  db.DB.QueryRow(userPostPointQuery, claims.Uuid, postId).Scan(&userPointing); if qErr != nil {
		log.Printf("This is the query error: %v", qErr)
    render.Render(w, r, ErrNotFound)
    return
  }

	insertPointQuery := `INSERT (UserId, PostId, Point) INTO "PointsOnPost" VALUES ($1, $2, $3)"`

	_, qErr = db.DB.Exec(insertPointQuery, claims.Uuid, postId, data.PointType); if qErr != nil {
		log.Printf("This is the query error: %v", qErr)
    render.Render(w, r, ErrNotFound)
    return
	}

  var points int
  getPointsQuery := `SELECT COUNT(*) FROM "PointsOnPost" WHERE UserId=$1`
  qErr = db.DB.QueryRow(getPointsQuery, claims.Uuid).Scan(&points); if qErr != nil {
		log.Printf("This is the query error: %v", qErr)
        render.Render(w, r, ErrNotFound)
        return
  }

  response := map[string]int {
    "points":points,
  }

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}

func RemovePointFromPostHandler(w http.ResponseWriter, r *http.Request) {

}

func mutatePostPointsHelper() {

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
