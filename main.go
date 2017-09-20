package main

import (
	"net/http"
	"log"
	"github.com/airIsland/webShow/route"
)

func main() {
	log.Println("start Server")
	http.Handle("/css/",http.FileServer(http.Dir("template")))
	http.Handle("/js/",http.FileServer(http.Dir("template")))
	http.Handle("/img/",http.FileServer(http.Dir("template")))

	http.HandleFunc("/login",route.LoginHandler)
	http.HandleFunc("/",route.NotFoundHander)
	http.ListenAndServe(":9009",nil)
	log.Println("Server started")
}
