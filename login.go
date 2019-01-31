package main

import (
	"net/http"
	"io/ioutil"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := User{}

	if r.Form.Get("id") != "" {
		id := r.Form.Get("id")
		pass := r.Form.Get("password")

		fileName := id + ".txt"
		body, err := ioutil.ReadFile(fileName)
		if err != nil {
			user = User{
				Id: id,
				Wrong : true,
			}
		}else{
			s := strings.Split(string(body), ",")
			if pass != s[1] {
				user = User{
					Id: s[0],
					Wrong : true,
				}
			}else {
				user = User{
					Id: s[0],
					Password: s[1],
					Name: s[2],
					Age: s[3],
					Country: s[4],
					Gender: s[5],
				}
				http.Redirect(w, r, "/home/?id="+s[0], http.StatusFound)
				return
			}
		}
	}
	render(w,"login.html", user)
}

