package main

import (
	"log"
	"net/http"
	"text/template"
)

//Globals
var tpl *template.Template

//Parsing initial templates removes need to parse at response function.
func init() {
	tpl = template.Must(template.ParseGlob("views/templates/*.gohtml"))
}

//Router
func main() {
	http.HandleFunc("/register/", register)
	http.HandleFunc("/login/", login)

	//Static Fileserver
	//http.Handle("login/assets/css/",
	//http.StripPrefix("/login/assets/css", http.FileServer(http.Dir("./views"))))

	http.ListenAndServe(":8080", nil)
}

//Login page handler
func login(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "login.gohtml", nil)
	if err != nil {
		log.Fatalln(w, http.StatusInternalServerError, err)
	}
}

//Register page handler
func register(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "register.gohtml", nil)
	if err != nil {
		log.Fatalln(w, http.StatusInternalServerError, err)
	}
}
