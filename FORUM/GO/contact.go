package main

import (
	"log"
	"net/http"
	"text/template"
)

type Contact struct {
	Id_contact string
	Firstname  string
	Lastname   string
	Email      string
	Message    string
	Date       string
}

//---------------affichage register---------------
func handler_contact(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../HTML/contact.html"))
	tmpl.Execute(w, nil)
}

//si utile met une demande dans la table
func handler_contact_exec(w http.ResponseWriter, r *http.Request) {
	contact := Contact{
		Firstname: r.FormValue("firstname"),
		Lastname:  r.FormValue("lastname"),
		Email:     r.FormValue("email"),
		Message:   r.FormValue("message"),
	}

	var err error
	_, err = db.Exec("INSERT INTO contact(id_contact, firstname, lastname, email, message, date) VALUES (DEFAULT,?,?,?,?,DATE(NOW()))", contact.Firstname, contact.Lastname, contact.Email, contact.Message)
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}
