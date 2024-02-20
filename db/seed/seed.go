package main

import (
	"context"
	"database/sql"
	"github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

var psqlConn string = "postgres://knostash_user:example@127.0.0.1:5432/knostash_db?sslmode=disable"
var loremIpsumPostFileName string = "LoremIpsumPostBody.txt"

func main() {
	fileContent, fileRdErr := ioutil.ReadFile(loremIpsumPostFileName)
	if fileRdErr != nil {
		log.Printf("File failed to be read with error: %v", fileRdErr)
	}
	db, sqlOpenErr := sql.Open("postgres", psqlConn)
	if sqlOpenErr != nil {
		log.Printf("SQL Connection error: %v", sqlOpenErr)
	}

	nanoId, _ := gonanoid.New()
	userOne := createUser(db, "Nikola_Tesla", "test", "nikola_tesla@example.com", nanoId)
	nanoId, _ = gonanoid.New()
	userTwo := createUser(db, "test", "test", "test@example.com", nanoId)

	for i := 0; i < 15; i++ {
		createPost(db, userOne, "Lorem Ipsum", string(fileContent))
	}
	lastPostId := createPost(db, userTwo, "More Lorem Ipsum", string(fileContent))

	for i := 0; i < 15; i++ {
		createComment(db, userOne, lastPostId, string(fileContent))
	}
}

func createUser(db *sql.DB, username string, password string, email string, uuid string) int {
	var lastId int

	query := `INSERT INTO "UserAccount" (Username, Password, Email, Active, Uuid) VALUES ($1, $2, $3, TRUE, $4) RETURNING ID`

	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashErr != nil {
		log.Fatal("Password hash error has occurred: ", hashErr)
	}

	sqlErr := db.QueryRow(query, username, hashedPassword, email, uuid).Scan(&lastId)
	if sqlErr != nil {
		log.Printf("SQL Error had occurred at insert: %v", sqlErr)
	}

	return lastId
}

func createPost(db *sql.DB, userId int, title string, body string) int {
	var lastPostId int

	query := `INSERT INTO
        "UserPost" (Owner, Title, Body)
        VALUES ($1, $2, $3)
		RETURNING ID
		`

	queryErr := db.QueryRow(query, userId, title, body).Scan(&lastPostId)
	if queryErr != nil {
		log.Printf("SQL Error had occurred at insert: %v", queryErr)
	}

	return lastPostId
}

func createComment(db *sql.DB, userId int, postId int, body string) int {
	var lastCommentId int

	ctx := context.Background()
	tx, txErr := db.BeginTx(ctx, nil)

	query := `INSERT INTO "Comment" (Owner, Body)
		VALUES ($1, $2)
		RETURNING ID`

	qErr := tx.QueryRowContext(ctx, query, userId, body).Scan(&lastCommentId)
	if qErr != nil {
		log.Printf("This is the query error: %v", qErr)
	}

	query = `INSERT INTO "CommentOnPost" (CommentId, PostId)
  VALUES ($1, $2)`

	_, qErr = tx.ExecContext(ctx, query, lastCommentId, postId)
	if qErr != nil {
		log.Printf("This is the query error: %v", qErr)
	}

	if txErr = tx.Commit(); txErr != nil {
		log.Printf("CreateComment SQL transaction error: %v", txErr)
	}

	return lastCommentId
}
