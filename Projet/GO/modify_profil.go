package main

import (
	"fmt"
	"net/http"
	"text/template"
)

//---------------modifier profil---------------
func handler_modifyprofil(w http.ResponseWriter, r *http.Request) {
	if verif(w, r) == true {

		tmpl2 := template.Must(template.ParseFiles("../HTML/settings.html"))

		modify_profil := info_user(w, r)

		fmt.Println(modify_profil)
		tmpl2.Execute(w, modify_profil)
	} else {
		http.Redirect(w, r, "/accueil#modal-opened", http.StatusFound)
	}
}

//---------------récup profil---------------
func handler_modifyprofil_exec(w http.ResponseWriter, r *http.Request) {
	if verif(w, r) == true && r.FormValue("email") != "" {
		var err error

		info_mofify := info_user(w, r)

		modify := User{
			Pseudo: r.FormValue("pseudo"),
			Email:  r.FormValue("email"),
		}

		_, err = db.Exec("UPDATE user SET email=?, pseudo=? WHERE email=? AND pseudo=?;", modify.Email, modify.Pseudo, info_mofify.Email, info_mofify.Pseudo)
		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/profil", http.StatusFound)
	} else if verif(w, r) == false {
		http.Redirect(w, r, "/accueil#modal-opened", http.StatusFound)
	} else if verif(w, r) == true && r.FormValue("email") == "" {
		http.Redirect(w, r, "/modifyprofil", http.StatusFound)
	}
}
