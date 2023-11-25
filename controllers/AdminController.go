package controllers

import (
  "net/http"
  "html/template"
  "log"
	"github.com/go-chi/render"
)

func AdminIndex(w http.ResponseWriter, r *http.Request) {
  var tmplFile = "static/admin.html"
  var navbarFile = "static/navbar.html"
  var textCard = "static/textcard.html"
  responseData := map[string]interface{} {
    "message": "Something bad happened",
  }

  tmpl, tmplErr := template.New("admin").ParseFiles(tmplFile, textCard, navbarFile)
  if tmplErr != nil {
    log.Printf("Tmpl Error Occurred: %v", tmplErr)
    render.Status(r, http.StatusInternalServerError)
    render.JSON(w, r, responseData)
    return
  }
  data := map[string]interface{}{
    "Name": "Jonathan Nguyen",
  }
  execErr := tmpl.Execute(w, data)
  if execErr != nil {
    log.Printf("tmpl exec went wrong: %v", execErr)
    render.Status(r, http.StatusInternalServerError)
    render.JSON(w, r, responseData)
    return
  }

  render.Status(r, http.StatusOK)
}

func AdminText(w http.ResponseWriter, r *http.Request) {
  var textCard = "static/textcard.html"
  responseData := map[string]interface{} {
    "message": "Something bad happened",
  }

  tmpl, tmplErr := template.New("textcard").ParseFiles(textCard)
  if tmplErr != nil {
    log.Printf("Tmpl Error Occurred: %v", tmplErr)
    render.Status(r, http.StatusInternalServerError)
    render.JSON(w, r, responseData)
    return
  }
  data := map[string]interface{}{
    "ChangeData": "Card with changed data",
  }
  execErr := tmpl.Execute(w, data)
  if execErr != nil {
    log.Printf("tmpl exec went wrong: %v", execErr)
    render.Status(r, http.StatusInternalServerError)
    render.JSON(w, r, responseData)
    return
  }
 
  render.Status(r, http.StatusOK)
}

