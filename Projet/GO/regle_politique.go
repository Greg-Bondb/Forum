package main

import (
	"net/http"
	"text/template"
)

func handler_reglement(w http.ResponseWriter, r *http.Request) {
	tmpl2 := template.Must(template.ParseFiles("../HTML/rules.html"))
	log := User{}
	if verif(w, r) {
		log = User{
			Pseudo: info_user(w, r).Pseudo,
			Error:  false,
			Error2: false,
		}
	} else {
		log.Pseudo = ""
		log.Error = true
		log.Error2 = err2
	}
	tmpl2.Execute(w, log)
	err2 = false
}

func handler_politique(w http.ResponseWriter, r *http.Request) {
	tmpl2 := template.Must(template.ParseFiles("../HTML/confidentiality.html"))
	log := User{}
	if verif(w, r) {
		log = User{
			Pseudo: info_user(w, r).Pseudo,
			Error:  false,
			Error2: false,
		}
	} else {
		log.Pseudo = ""
		log.Error = true
		log.Error2 = err2
	}
	tmpl2.Execute(w, log)
	err2 = false
}
