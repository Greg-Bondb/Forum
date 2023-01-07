package main

import (
	"net/http"
	"text/template"
)

func handler_modifypassword(w http.ResponseWriter, r *http.Request) {
	if verif(w, r) == true {

		tmpl2 := template.Must(template.ParseFiles("../HTML/security.html"))

		info_mofify := info_user(w, r)

		if verif(w, r) {
			info_mofify = User{
				Pseudo: info_user(w, r).Pseudo,
				Error:  false,
				Error2: false,
			}
		} else {
			info_mofify.Pseudo = ""
			info_mofify.Error = true
			info_mofify.Error2 = err2
		}

		tmpl2.Execute(w, info_mofify)
	} else {
		http.Redirect(w, r, "/accueil", http.StatusFound)
	}
}

func handler_modifypassword_exec(w http.ResponseWriter, r *http.Request) {
	if verif(w, r) == true && r.FormValue("password") != "" {
		var err error

		info_mofify := info_user(w, r)

		if r.FormValue("password") == r.FormValue("password2") {

			modify := User{
				Password: r.FormValue("password"),
			}

			_, err = db.Exec("UPDATE `user` SET `password`= MD5(?) WHERE email = ? AND pseudo=?;", modify.Password, info_mofify.Email, info_mofify.Pseudo)
			if err != nil {
				panic(err)
			}

			http.Redirect(w, r, "/profil", http.StatusFound)
		} else {
			http.Redirect(w, r, "/security", http.StatusFound)
		}
	} else if verif(w, r) == false {
		http.Redirect(w, r, "/accueil", http.StatusFound)
	} else if verif(w, r) == true && r.FormValue("password") == "" {
		http.Redirect(w, r, "/security", http.StatusFound)
	}
}
