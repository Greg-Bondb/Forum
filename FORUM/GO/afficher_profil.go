package main

import (
	"log"
	"net/http"
	"text/template"
)

//Structure qui stock toute la liste des users
type TabUser struct {
	User []User
}

//---------------affichage profil---------------

func handler_profil(w http.ResponseWriter, r *http.Request) {
	if verif(w, r) {
		tmpl2 := template.Must(template.ParseFiles("../HTML/details-of-account.html"))
		data := TabUser{}

		infoUser, err := db.Query("SELECT id_user, id_role, firstname, lastname, pseudo, born, email, password, lang, creation FROM user;")
		if err != nil {
			log.Fatal(err)
		}

		for infoUser.Next() {
			dbUser := User{}

			err := infoUser.Scan(&dbUser.ID_user, &dbUser.ID_role, &dbUser.Firstname, &dbUser.Lastname, &dbUser.Pseudo, &dbUser.Born, &dbUser.Email, &dbUser.Password, &dbUser.Lang, &dbUser.Creation)
			if err != nil {
				log.Fatal(err)
			}

			log.Println(dbUser.ID_user, dbUser.ID_role, dbUser.Firstname, dbUser.Lastname, dbUser.Pseudo, dbUser.Born, dbUser.Email, dbUser.Password, dbUser.Lang)
			data.User = append(data.User, dbUser)
		}
		defer infoUser.Close()

		verif(w, r)
		tmpl2.Execute(w, data.User[profil(data, w, r)])
	} else {
		http.Redirect(w, r, "/accueil#modal-opened", http.StatusFound)
	}
}
