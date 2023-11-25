package porkmiddleware

import (
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
  "os"
)

// middleware

var jwtSecret []byte = []byte(os.Getenv("JWT_SECRET"))

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookieValue, cookieErr := r.Cookie("jwtCookie")
		if cookieErr != nil {
			render.Render(w, r, ErrNotFound)
			return
		}

		token, err := jwt.ParseWithClaims(cookieValue.Value, &KnoAuthCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			log.Printf("This is the error: %v", err)
			render.Render(w, r, ErrNotFound)
			return
		}
		if claims, ok := token.Claims.(*KnoAuthCustomClaims); ok && token.Valid {
			log.Printf("%v %v", claims.Username, claims.RegisteredClaims.Issuer)
		} else {
			log.Println(err)
			render.Render(w, r, ErrNotFound)
			return
		}

		// Serve the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

type KnoAuthCustomClaims struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found"}
