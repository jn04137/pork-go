package controllers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strconv"
	"encoding/json"

	"porkgo/db"
	"porkgo/models"
	"porkgo/porkmiddleware"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	postId, parseErr := strconv.Atoi(chi.URLParam(r, "postId"))
	if parseErr != nil {
		log.Printf("parse error")
	}
	cookie, cookieErr := r.Cookie("jwtCookie")
	if cookieErr != nil {
		log.Printf("This is the cookieErr: %v", cookieErr)
		render.Render(w, r, ErrNotFound)
		return
	}

	data := &models.CommentReq{}
	if bindErr := render.Bind(r, data); bindErr != nil {
		log.Printf("Binding error: %v", bindErr)
		render.Render(w, r, ErrNotFound)
		return
	}

	// Need token to create the comment with correct user as owner
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

	ctx := context.Background()

	tx, err := db.DB.BeginTx(ctx, nil)

	// Insert into the UserComment Table
	var commentId int
	query := `INSERT INTO "Comment" (Owner, Body)
		VALUES ((SELECT ID FROM "UserAccount" WHERE Uuid=$1), $2)
		RETURNING ID`
	qErr := tx.QueryRowContext(ctx, query, claims.Uuid, data.Body).Scan(&commentId)
	if qErr != nil {
		log.Printf("This is the query error: %v", qErr)
		render.Render(w, r, ErrNotFound)
		return
	}

	// Insert the relation into the junction table
	query = `INSERT INTO "CommentOnPost" (CommentId, PostId)
		VALUES ($1, $2)`
	_, qErr = tx.ExecContext(ctx, query, commentId, postId)
	if qErr != nil {
		log.Printf("This is the query error: %v", qErr)
		render.Render(w, r, ErrNotFound)
		return
	}

	if err = tx.Commit(); err != nil {
		log.Printf("SQL transaction error: %v", err)
	}

	responseData := map[string]interface{}{
		"message": "was a success",
	}
	
	render.Status(r, http.StatusOK)
	render.JSON(w, r, responseData)
}

func LoadCommentsHandler(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "postId")
	commentCursor, cursorErr := strconv.Atoi(chi.URLParam(r, "commentCursor"))
	if cursorErr != nil {
		log.Printf("There was an error reading the commentCursor")
	}

	query := `SELECT ua.Username, c.Id, c.Body, c.CreatedAt, c.IsDeleted FROM "CommentOnPost" as cop
		INNER JOIN "Comment" as c ON cop.CommentId=c.ID
		INNER JOIN "UserAccount" as ua ON ua.ID=c.Owner
		WHERE cop.PostId=$1
    ORDER BY c.CreatedAt DESC
		LIMIT 5
		OFFSET $2
		`
	rows, qErr := db.DB.Query(query, postId, commentCursor)
	if qErr != nil {
		log.Printf("sql query error: %v", qErr)
	}

	var listOfComments []models.CommentModel = make([]models.CommentModel, 0)
	for rows.Next() {
		var commentScans models.CommentModel
		if rowsErr := rows.Scan(
			&commentScans.Owner,
			&commentScans.CommentId,
			&commentScans.Body,
			&commentScans.CreatedAt,
			&commentScans.IsDeleted); rowsErr != nil {
			log.Printf("Rows scan error: %v", rowsErr)
			continue
		}
		isCommentDeleted(&commentScans)
		listOfComments = append(listOfComments, commentScans)
	}
	responseData := map[string]interface{}{
		"comments":   listOfComments,
		"nextCursor": commentCursor + 5,
	}
	if len(listOfComments) == 0 {
		responseData["nextCursor"] = nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
	return
}

/*
	Deleting comment should just change the isDeleted flag to true 
	on the comment
*/
func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	jwtUserInfo := r.Context().Value("jwtUserValues").(porkmiddleware.JwtUserValues)
	var commentModel models.CommentModel
	err := json.NewDecoder(r.Body).Decode(&commentModel)
	if err != nil {
		log.Printf("This is the error: %v", err)
		render.Render(w, r, ErrNotFound)
		return
	}

	query := `UPDATE "Comment" SET isDeleted=TRUE 
		WHERE ID=$1 and Owner=(SELECT ID FROM "UserAccount" WHERE Uuid=$2)`
	
	responseData := map[string]interface{}{
		"message": "success",
	}
	
	_, qErr := db.DB.Exec(query, commentModel.CommentId, jwtUserInfo.Uuid); if qErr != nil {
		log.Printf("This is the query error: %v", qErr)
		render.Render(w, r, ErrNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
	return
}

/*
	Should allow the user to change the body of the comment and update the
	lastUpdated field
*/
func EditCommentHandler(w http.ResponseWriter, r *http.Request) {

	responseData := map[string]interface{}{
		"message": "success",
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
	return
}

// Checks if comment has deleted flag
// comment prints as if deleted if flag is true
func isCommentDeleted(comment *models.CommentModel) *models.CommentModel {
	if comment.IsDeleted {
		comment.Body = "Deleted Comment"
		comment.Owner = "Anonymous User"
	}

	return comment
}

