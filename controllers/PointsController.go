package controllers

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"

	"porkgo/db"
)

func GetPostPoints(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "postId")
	var postPoints string
	var userPostPointStatus string
	response := make(map[string]string)

	getPointsQuery := `SELECT 
    (SELECT COUNT(*)
    FROM "PointsOnPost"
    WHERE "PointsOnPost".PostId=$1
    AND "PointsOnPost".Point='plus') - 
    (SELECT COUNT(*)
    FROM "PointsOnPost"
    WHERE "PointsOnPost".PostId=$1
    AND "PointsOnPost".Point='minus')`
	qErr := db.DB.QueryRow(getPointsQuery, postId).Scan(&postPoints)
	if qErr != nil {
		render.Render(w, r, ErrNotFound)
		return
	}
	response["points"] = postPoints

	cookie, cookieErr := r.Cookie("jwtCookie")
	if cookieErr != nil {
		response["userPostPointStatus"] = "empty"
		render.Status(r, http.StatusOK)
		render.JSON(w, r, response)
		return
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &KnoAuthCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		log.Printf("This is the error: %v", err)
	}
	claims, ok := token.Claims.(*KnoAuthCustomClaims)

	getUserLikeStatusOnPost := `SELECT Point FROM "PointsOnPost" 
    WHERE "PointsOnPost".UserId=(SELECT ID FROM "UserAccount" WHERE "UserAccount".Uuid=$1) 
    AND "PointsOnPost".PostId=$2`
	if token.Valid && ok {
		qErr = db.DB.QueryRow(getUserLikeStatusOnPost, claims.Uuid, postId).Scan(&userPostPointStatus)
		if qErr == sql.ErrNoRows {
			userPostPointStatus = "empty"
		} else if qErr != nil {
			log.Printf("(line 32) This is the query error: %v", qErr)
			render.Render(w, r, ErrNotFound)
			return
		}
	}

	response["userPostPointStatus"] = userPostPointStatus

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}

// TODO This controller needs to be cleaned up.
// The minus method is going to be a mess too
func MutatePointToPostHandler(w http.ResponseWriter, r *http.Request) {
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

	response, pointsMutateErr := mutatePostPointsHelper(data.Pointing, claims.Uuid, postId)
	if pointsMutateErr != nil {
		log.Printf("mutation error: %v", pointsMutateErr)
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, ErrNotFound)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
	return
}

func RemovePointFromPostHandler(w http.ResponseWriter, r *http.Request) {

}

func AddPointFromCommentHandler(w http.ResponseWriter, r *http.Request) {

}

func RemovePointFromCommentHandler(w http.ResponseWriter, r *http.Request) {

}

// pointAction should either be 'add' or 'remove'
/*
  Needs to respond with the current points on the post after mutation
  Should probably use a transaction to get more consistent information
*/
func mutatePostPointsHelper(pointAction string, uuid string, postId string) (map[string]string, error) {

	responseData := make(map[string]string)
	var points string
	var userPointing string

	userPostPointQuery := `SELECT Point
		FROM "PointsOnPost"
		WHERE "PointsOnPost".UserId=(SELECT ID from "UserAccount" WHERE Uuid=$1)
		AND "PointsOnPost".PostId=$2`
	getPointsQuery := `SELECT 
    (SELECT COUNT(*)
    FROM "PointsOnPost"
    WHERE "PointsOnPost".PostId=$1
    AND "PointsOnPost".Point='plus') - 
    (SELECT COUNT(*)
    FROM "PointsOnPost"
    WHERE "PointsOnPost".PostId=$1
    AND "PointsOnPost".Point='minus')`
	qErr := db.DB.QueryRow(userPostPointQuery, uuid, postId).Scan(&userPointing)
	if qErr == sql.ErrNoRows {
		// If record doesn't exist, insert a new record with a like
		insertPointQuery := `INSERT INTO "PointsOnPost" (UserId, PostId, Point)
			VALUES ((SELECT ID FROM "UserAccount" WHERE "UserAccount".Uuid=$1), $2, $3)`
		_, qErr = db.DB.Exec(insertPointQuery, uuid, postId, pointAction)
		if qErr != nil {
			return responseData, qErr
		}

		// Query for the total number of points on post and return to client
		qErr = db.DB.QueryRow(getPointsQuery, postId).Scan(&points)
		if qErr != nil {
			return responseData, qErr
		}
		responseData["points"] = points
		return responseData, qErr
	} else if qErr != nil {
		return responseData, qErr
	}

	updatePointsQuery := `UPDATE "PointsOnPost" 
    SET Point=$1 
    WHERE UserId=(SELECT ID FROM "UserAccount" WHERE "UserAccount".Uuid=$2) AND PostId=$3`
	if userPointing == pointAction {
		_, qErr = db.DB.Exec(updatePointsQuery, "empty", uuid, postId)
		if qErr != nil {
			return responseData, qErr
		}
	} else {
		_, qErr = db.DB.Exec(updatePointsQuery, pointAction, uuid, postId)
		if qErr != nil {
			return responseData, qErr
		}
	}

	qErr = db.DB.QueryRow(getPointsQuery, postId).Scan(&points)
	if qErr != nil {
		return responseData, qErr
	}

	var userPostPointStatus string
	getUserLikeStatusOnPost := `SELECT Point FROM "PointsOnPost" 
    WHERE "PointsOnPost".UserId=(SELECT ID FROM "UserAccount" WHERE "UserAccount".Uuid=$1) 
    AND "PointsOnPost".PostId=$2`
	qErr = db.DB.QueryRow(getUserLikeStatusOnPost, uuid, postId).Scan(&userPostPointStatus)
	if qErr != nil {
		log.Printf("(line 32) This is the query error: %v", qErr)
		return responseData, qErr
	}

	responseData["points"] = points
	responseData["userPostPointStatus"] = userPostPointStatus
	return responseData, qErr
}

type Point struct {
	Pointing string `json:"pointing"`
}

type PointReq struct {
	*Point
}

func (a *PointReq) Bind(r *http.Request) error {
	return nil
}
