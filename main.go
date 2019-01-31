package main

import (
	"net/http"
	"html/template"
	"log"
)

type User struct {
	Id string
	Password string
	Name string
	Gender string
	Age string
	Country string
	Wrong bool
}

type File struct {
	Title string
	Body []byte
}

func main() {
	http.HandleFunc("/home/", Home)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/saveRegister", SaveRegister)
	// http.HandleFunc("/afterLogin", AfterLogin)
	log.Fatal(http.ListenAndServe(":80",nil))
}

func render(w http.ResponseWriter, tmpl string, user User) {
	t, err := template.ParseFiles(tmpl)

	if err != nil {
		log.Print("error parse : ",err)
	}

	err = t.Execute(w, user)
	
	if err != nil {
		log.Print("error execute : ",err)
	}
}
