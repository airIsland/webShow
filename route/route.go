package route

import (
	"net/http"
	"log"
	"html/template"
	"strings"
	"reflect"
	"github.com/airIsland/webShow/controller"
)

func LoginHandler(w http.ResponseWriter,r *http.Request){
	log.Println("loginHandler")
	    pathInfo := strings.Trim(r.URL.Path, "/")
	    parts := strings.Split(pathInfo, "/")
	    var action = ""
	    if len(parts) > 1 {
	        action = strings.Title(parts[1]) + "Action"
	    }

	    login := &controller.LoginController{}
	    controller := reflect.ValueOf(login)
	    method := controller.MethodByName(action)
	    if !method.IsValid() {
	        method = controller.MethodByName(strings.Title("index") + "Action")
	    }
	    requestValue := reflect.ValueOf(r)
	    responseValue := reflect.ValueOf(w)
	    method.Call([]reflect.Value{responseValue, requestValue})
}


func NotFoundHander(w http.ResponseWriter, r *http.Request){
	if r.URL.Path == "/"{
		http.Redirect(w,r,"/login",http.StatusFound)
	}

	t,err :=template.ParseFiles("template/html/404.html")
	if err!=nil {
		log.Println(err)
	}
	t.Execute(w,nil)
}