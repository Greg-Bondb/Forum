package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

//---------------affichage login---------------

func handler_login(w http.ResponseWriter, r *http.Request) {
	tmpl2 := template.Must(template.ParseFiles("../HTML/homepage.html"))
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

//---------------récup login---------------

func handler_login_exec(w http.ResponseWriter, r *http.Request) {
	if verif(w, r) == false && r.FormValue("email") != "" {

		login := User{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		login.Password = crypt(login.Password)

		rows, err := db.Query("SELECT id_user, id_role, firstname, lastname, pseudo, born, email, password, lang,creation  FROM user WHERE email=? AND password=?", login.Email, login.Password)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			if err := rows.Scan(&login.ID_user, &login.ID_role, &login.Firstname, &login.Lastname, &login.Pseudo, &login.Born, &email, &password, &login.Lang, &login.Creation); err != nil {
				log.Fatal(err)
			}
		}

		if login.Email == email && login.Password == password {
			fmt.Println(login.Email, email, login.Password, password)

			fmt.Println("Connexion")

			cookie(login.ID_user, w, r)

			http.Redirect(w, r, "/accueil", http.StatusFound)

		} else if login.Email != email || login.Password != password {

			fmt.Println("Connexion refusée")
			err2 = true
			http.Redirect(w, r, "/accueil#modal-opened", http.StatusFound)

		}
	} else if verif(w, r) == true {

		http.Redirect(w, r, "/accueil", http.StatusFound)
	} else if verif(w, r) == false && r.FormValue("email") == "" {
		http.Redirect(w, r, "/accueil", http.StatusFound)
	}
}

func handler_login_delete(w http.ResponseWriter, r *http.Request) {
	var cookie, err = r.Cookie("ID_user")
	fmt.Println(err)
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/accueil", http.StatusFound)
}
