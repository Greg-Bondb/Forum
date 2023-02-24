package main

import (
	"log"
	"net/http"
	"text/template"
)

//Structure Creaforum
type Creaforum struct {
	ID_user       string
	Title_subject string
	Categories    string
	Types         string
	Messages      string
	Date          string
	Error         bool
}

//---------------affichage créa forum---------------
func handler_creationforum(w http.ResponseWriter, r *http.Request) {
	if verif(w, r) == true {
		tmpl3 := template.Must(template.ParseFiles("../HTML/creation_forum.html"))
		tmpl3.Execute(w, nil)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

}

//---------------récup créa forum---------------
func handler_creationforum_exec(w http.ResponseWriter, r *http.Request) {
	if verif(w, r) == true && r.FormValue("title_subject") != "" {
		information := Creaforum{
			Title_subject: r.FormValue("title_subject"),
			Categories:    r.FormValue("categories"),
			Types:         r.FormValue("types"),
			Messages:      r.FormValue("message"),
			ID_user:       return_cookies(w, r),
		}

		var err error
		_, err = db.Exec("INSERT INTO tabforum(id_forum, title_subject, categories, types, messages, date, id_user) VALUES (DEFAULT,?,?,?,?,DATE(NOW()),?)", information.Title_subject, information.Categories, information.Types, information.Messages, information.ID_user)
		if err != nil {
			log.Fatal(err)
		}

		http.Redirect(w, r, "/subject_affichage", http.StatusFound)
	} else if verif(w, r) == false {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else if verif(w, r) == true && r.FormValue("title_subject") == "" {
		http.Redirect(w, r, "/subject", http.StatusFound)
	}
}
