package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	_ "github.com/lib/pq"
	"log"
	"net/http"
  "mime"

	"banhbao/porkgo/controllers"
	"banhbao/porkgo/db"
	"banhbao/porkgo/porkmiddleware"
)

func main() {
	if err := db.Open(); err != nil {
		log.Fatal("DB Connection issue")
	}
	defer db.Close()

	r := chi.NewRouter()
  mimeErr1 := mime.AddExtensionType(".js", "text/javascript"); if mimeErr1 != nil {
    log.Printf("error: %v", mimeErr1)
  }
  mimeErr2 := mime.AddExtensionType(".css", "text/css"); if mimeErr2 != nil {
    log.Printf("error: %v", mimeErr2)
  }
	log.Printf("Application has started")
	r.Use(middleware.Logger)
  

	defer db.DB.Close()

  r.Use(middleware.AllowContentType("application/json", "text/css"))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}))
  
  r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
    responseData := map[string]interface{}{
      "message": "pong",
    }

    render.Status(r, http.StatusOK)
    render.JSON(w, r, responseData)
  })

	r.Route("/api/", func(r chi.Router) {
    r.Route("/protected", func(r chi.Router) {
      r.Use(porkmiddleware.JWTAuthMiddleware)
      r.Get("/checktoken", func(w http.ResponseWriter, r *http.Request) {
        data := &SomeJSONMsgResp{
          Message: "This is the message",
        }

        render.Status(r, http.StatusOK)
        render.JSON(w, r, data)
      })
      r.Post("/createpost", controllers.CreatePostHandler)
      r.Post("/createcomment/{postId}", controllers.CreateCommentHandler)
      r.Post("/post/addpoints/{postId}", controllers.AddPointToPostHandler)
      r.Post("/post/removepoints/{postId}", controllers.RemovePointFromPostHandler)
    })
    r.Route("/public", func(r chi.Router) {
      r.Get("/getfeedposts/{page}", controllers.LoadFeedPost)
      r.Get("/viewpost/{postId}", controllers.ViewPost)
      r.Get("/loadcomments/{postId}/{commentCursor}", controllers.LoadCommentsHandler)
      r.Get("/post/points/{postId}", controllers.GetPostPoints)
      r.Post("/signup", controllers.SignupHandler)
      r.Post("/login", controllers.LoginHandler)
      r.Get("/isloggedin", controllers.IsLoggedInHandler)
      r.Get("/AuthEmailTest", controllers.AuthEmail)
      r.Get("/verifyaccount/{token}", controllers.VerifyAccount)
      r.Get("/resendaccountverificationtoken", controllers.VerifyAccount)
      r.Route("/admin", func(r chi.Router) {
      r.Get("/", controllers.AdminIndex)
      r.Get("/test", controllers.AdminText)
	})
    })
	})
	

	http.ListenAndServe(":8000", r)
}

type SomeJSONMsg struct {
	Message string `json:"message"`
}

type SomeJSONMsgResp struct {
	Message string `json:"message"`
}
