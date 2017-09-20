package controller

import (
	"net/http"
	"log"
	"html/template"
)

type LoginController struct {
}

func (this *LoginController)IndexAction(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("template/html/index.html")
    if (err != nil) {
        log.Println(err)
    }
    t.Execute(w, nil)
}