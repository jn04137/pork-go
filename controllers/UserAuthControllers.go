package controllers

import (
	"porkgo/db"
	"porkgo/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
	"github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"time"
)

var smtpPass string = os.Getenv("SMTP_PASS")
var backendUrl string = os.Getenv("KNOSTASH_BACKEND_URL")
var jwtSecret []byte = []byte(os.Getenv("JWT_SECRET"))

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	data := &models.UserAuthReq{}

	if err := render.Bind(r, data); err != nil {
		log.Printf("Something went wrong")
		return
	}

	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if hashErr != nil {
		log.Printf("Hash went wrong")
		render.Render(w, r, ErrNotFound)
		return
	}

	ctx := context.Background()
	tx, err := db.DB.BeginTx(ctx, nil)
	var uuid string
	nanoId, nanoIdErr := gonanoid.New()
	if nanoIdErr != nil {
		log.Printf("nanoId error: %v", nanoIdErr)
		render.Render(w, r, ErrNotFound)
		return
	}
	createUserQuery := "INSERT INTO \"UserAccount\" (Username, Password, Email, Uuid) VALUES ($1, $2, $3, $4) RETURNING Uuid"
	qErr := tx.QueryRowContext(ctx, createUserQuery, data.Username, hashedPassword, data.Email, nanoId).Scan(&uuid)
	if qErr != nil {
		log.Printf("create user: %v", err)
		render.Render(w, r, ErrNotFound)
		return
	}

	if err = tx.Commit(); err != nil {
		log.Printf("tx error: %v", err)
		render.Render(w, r, ErrNotFound)
		return
	}

	claims := EmailVerificationClaims{
		data.Email,
		uuid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, tokenErr := token.SignedString(jwtSecret)
	if tokenErr != nil {
		log.Printf("Token err: %v", tokenErr)
		render.Render(w, r, ErrNotFound)
		return
	}

	// Contents of the message
	fromMsg := "From: thdr.site@gmail.com\n"
	toMsg := fmt.Sprintf("To: %v\n", data.Email)
	subject := "Subject: Test email from Go!\n"
	mime := "MIME-version: 1.0; \nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf("<html><body>Verify your account at this link: %v/api/public/verifyaccount/%v</body></html>", backendUrl, ss)
	msg := []byte(fromMsg + toMsg + subject + mime + body)

	// Configuration for sending an email
	from := "thdr.site@gmail.com"
	pass := smtpPass
	to := data.Email
	emailErr := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if emailErr != nil {
		log.Printf("smtp error: %v", err)
		render.Render(w, r, ErrNotFound)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, successfulResponse)
}

func AuthEmail(w http.ResponseWriter, r *http.Request) {

	// Contents of the message
	fromMsg := "From: thdr.site@gmail.com\n"
	toMsg := "To: jnganguyen3@gmail.com\n"
	subject := "Subject: Test email from Go!\n"
	mime := "MIME-version: 1.0; \nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := "<html><body>Verify your account</body></html>"
	msg := []byte(fromMsg + toMsg + subject + mime + body)

	// Configuration for sending an email
	from := "thdr.site@gmail.com"
	pass := smtpPass
	to := "jnganguyen3@gmail.com"
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		render.Render(w, r, ErrNotFound)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, successfulResponse)
	return
}

func VerifyAccount(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	query := `UPDATE "UserAccount" SET Active=true WHERE Uuid=$1;`

	parsedToken, tokenParseErr := jwt.ParseWithClaims(
		token,
		&EmailVerificationClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
	if tokenParseErr != nil {
		log.Printf("tokenParseErr: %v", tokenParseErr)
		render.Render(w, r, ErrNotFound)
		return
	} else if claims, ok := parsedToken.Claims.(EmailVerificationClaims); ok {
		log.Printf("Claims info: %v", claims.Uuid)
		_, qErr := db.DB.Exec(query, claims.Uuid)
		if qErr != nil {
			log.Printf("This is the query error: %v", qErr)
			render.Render(w, r, ErrNotFound)
			return
		}
	} else {
		log.Printf("claims not okay: %v", tokenParseErr)
		render.Render(w, r, ErrNotFound)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, successfulResponse)
}

func ResendAccountVerification(w http.ResponseWriter, r *http.Request) {

	render.Status(r, http.StatusOK)
	render.JSON(w, r, successfulResponse)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	data := &models.UserAuthReq{}
	var qUser models.UserModel

	if err := render.Bind(r, data); err != nil {
		log.Printf("Binding error")
		render.Render(w, r, ErrNotFound)
		return
	}

	// Find the existing user if exists, if err then pass fail code
	query := "SELECT username, password, uuid FROM \"UserAccount\" WHERE USERNAME=$1 AND Active=true"
	err := db.DB.QueryRow(query, data.Username).Scan(&qUser.Username, &qUser.Password, &qUser.Uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("User was not found")
			render.Render(w, r, ErrNotFound)
			return
		}
	}

	// Compare password to hash, if err then pass fail code
	err = bcrypt.CompareHashAndPassword([]byte(qUser.Password), []byte(data.Password))
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}

	// Generate JWT and store it in cookie for user
	var (
		t *jwt.Token
		s string
	)

	claims := KnoAuthCustomClaims{
		qUser.Uuid,
		qUser.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	//key = []byte("testhash")
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err = t.SignedString(jwtSecret)
	if err != nil {
		// TODO need to return 500 error later
		render.Render(w, r, ErrNotFound)
		return
	}

	cookie := &http.Cookie{
		Name:     "jwtCookie",
		Value:    s,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(w, cookie)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, successfulResponse)
}

func IsLoggedInHandler(w http.ResponseWriter, r *http.Request) {
	responseData := map[string]interface{}{
		"isLoggedIn": false,
		"username":   "",
	}
	cookieValue, cookieErr := r.Cookie("jwtCookie")
	if cookieErr != nil {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, responseData)
		return
	}

	token, err := jwt.ParseWithClaims(cookieValue.Value, &KnoAuthCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, responseData)
		return
	}
	claims, ok := token.Claims.(*KnoAuthCustomClaims)
	if !ok {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, responseData)
		return
	}

	responseData["isLoggedIn"] = true
	responseData["username"] = claims.Username

	render.Status(r, http.StatusOK)
	render.JSON(w, r, responseData)
}

var ErrNotFound = &models.ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found"}

var successfulResponse map[string]interface{} = map[string]interface{}{
	"message": "Success",
}

type KnoAuthCustomClaims struct {
	Uuid     	string `json:"uuid"`
	Username 	string `json:"username"`
	jwt.RegisteredClaims
}

type EmailVerificationClaims struct {
	Email string `json:"email"`
	Uuid  string `json:"uuid"`
	jwt.RegisteredClaims
}
