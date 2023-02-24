package main

import (
	"log"
	"net/http"
	"text/template"
)

//Strcuture User
type User struct {
	ID_user   string
	ID_role   string
	Firstname string
	Lastname  string
	Pseudo    string
	Born      string
	Email     string
	Password  string
	Lang      string
	Creation  string
	Error     bool
	Error2    bool
}

//couple la fonction affichage avec celle de recuperation de creation user

func creauser() {

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if verif(w, r) == false {

			tmpl := template.Must(template.ParseFiles("../HTML/creation_compte.html"))
			tmpl.Execute(w, nil)

		} else {

			http.Redirect(w, r, "/profil", http.StatusFound)

		}
	})

	http.HandleFunc("/register_exec", func(w http.ResponseWriter, r *http.Request) {

		if verif(w, r) == false && r.FormValue("email") != "" {

			register := User{
				Firstname: r.FormValue("firstname"),
				Lastname:  r.FormValue("lastname"),
				Pseudo:    r.FormValue("pseudo"),
				Born:      r.FormValue("born"),
				Email:     r.FormValue("email"),
				Password:  r.FormValue("password"),
				Lang:      r.FormValue("lang"),
			}

			_, err := db.Exec("INSERT INTO user (id_user, id_role, firstname, lastname, pseudo, born, email, password, lang, creation) VALUES (DEFAULT, DEFAULT, ?, ?, ?, ?, ?, MD5(?), ?, DATE(NOW()) )", register.Firstname, register.Lastname, register.Pseudo, register.Born, register.Email, register.Password, register.Lang)
			if err != nil {
				register.Error = true
			}

			//---------------------------

			infoUser, err := db.Query("SELECT id_user, id_role, firstname, lastname, pseudo, born, email, password, lang, creation FROM user WHERE email=?;", register.Email)
			if err != nil {
				log.Fatal(err)
			}
			dbUser := User{}
			for infoUser.Next() {

				err := infoUser.Scan(&dbUser.ID_user, &dbUser.ID_role, &dbUser.Firstname, &dbUser.Lastname, &dbUser.Pseudo, &dbUser.Born, &dbUser.Email, &dbUser.Password, &dbUser.Lang, &dbUser.Creation)
				if err != nil {
					log.Fatal(err)
				}
			}

			cookie(dbUser.ID_user, w, r)

			http.Redirect(w, r, "/profil", http.StatusFound)

		} else if verif(w, r) == true {

			http.Redirect(w, r, "/profil", http.StatusFound)

		} else if verif(w, r) == false && r.FormValue("email") == "" {

			http.Redirect(w, r, "/register", http.StatusFound)

		}
	})
}
