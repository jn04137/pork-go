package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	_ "github.com/lib/pq"
	"log"
	"net/http"
  "strings"
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
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
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

	r.Route("/public", func(r chi.Router) {
		r.Get("/getfeedposts/{page}", controllers.LoadFeedPost)
		r.Get("/viewpost/{postId}", controllers.ViewPost)
		r.Get("/loadcomments/{postId}/{commentCursor}", controllers.LoadCommentsHandler)
	})

	r.Route("/api", func(r chi.Router) {
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
	})

	// This is the route for auth based endpoints (login, signup, etc.)
	r.Route("/auth", func(r chi.Router) {
		r.Post("/signup", controllers.SignupHandler)
		r.Post("/login", controllers.LoginHandler)
		r.Get("/isloggedin", controllers.IsLoggedInHandler)
    r.Get("/AuthEmailTest", controllers.AuthEmail)
    r.Get("/verifyaccount/{token}", controllers.VerifyAccount)
    r.Get("/resendaccountverificationtoken", controllers.VerifyAccount)
	})
	
  r.Route("/admin", func(r chi.Router) {
    r.Get("/", controllers.AdminIndex)
    r.Get("/test", controllers.AdminText)
	})

  r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
  r.Handle("/assets/*", http.StripPrefix("/assets", http.FileServer(http.Dir("/dist/assets"))))

  fs := http.FileServer(http.Dir("/dist"))
  r.Handle("/", fs)
  r.Get("/*", func (w http.ResponseWriter, httpReq *http.Request) {
    http.ServeFile(w, httpReq, "/dist/index.html")
    return
  })

	http.ListenAndServe(":8000", r)
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		log.Printf("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

type SomeJSONMsg struct {
	Message string `json:"message"`
}

type SomeJSONMsgResp struct {
	Message string `json:"message"`
}
