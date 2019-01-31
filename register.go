package main

import (
	"net/http"
	"io/ioutil"
)

func Register(w http.ResponseWriter, r *http.Request) {
	user := User{}

	render(w,"register.html",user)
}

func SaveRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id := r.Form.Get("id")
	password := r.Form.Get("password")
	name := r.Form.Get("name")
	age := r.Form.Get("age")
	gender := r.Form.Get("gender")
	country := r.Form.Get("country")

	user := User {
		Id: id,
		Password: password,
		Name: name,
		Age: age,
		Country: country,
		Gender: gender,
	}

	user.save()

	http.Redirect(w, r, "/login", http.StatusFound)
}

func (user *User)save() error{
	fileName := user.Id + ".txt"
	body := []byte(user.Id+","+user.Password+","+user.Name+","+user.Age+
			","+user.Country+","+user.Gender+",")
	return ioutil.WriteFile(fileName, body, 0600)
}
