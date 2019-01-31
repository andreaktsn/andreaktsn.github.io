package main

import (
	"net/http"
	"io/ioutil"
	"strings"
	"log"
)

func Home(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")
	log.Print(id)
	user := User{}

	fileName := id + ".txt"
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		user = User{
			Wrong: true,
		}
		log.Print("file not found")
	}else{
		s := strings.Split(string(body), ",")
		user = User{
			Id: s[0],
			Password: s[1],
			Name: s[2],
			Age: s[3],
			Country: s[4],
			Gender: s[5],
		}	
	}
	render(w,"home.html",user)
}