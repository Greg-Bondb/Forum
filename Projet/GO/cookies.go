package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//---------------func cookie---------------
type Cookie struct {
	Name      string
	Value     string
	Path      string
	Domain    string
	Expires   time.Time
	RawExpire string
	MaxAge    int
	Secure    bool
	HttpOnly  bool
	Raw       string
	Unparsed  []string
}

//---------------envoie cookie---------------
func cookie(id string, w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "ID_user",
		Value:  id,
		MaxAge: 3000,
	}
	fmt.Print("set cookie")
	http.SetCookie(w, cookie)
}

//verifie si il y a un cookie
func verif(w http.ResponseWriter, req *http.Request) bool {
	var cookie, err = req.Cookie("ID_user")
	if err == nil {
		if cookie.Value != "" {
			return true
		}
		return false
	} else {
		fmt.Println("marchepas")
		return false
	}
}

//sert a retourner id du cookies
func return_cookies(w http.ResponseWriter, req *http.Request) string {
	var cookie, err = req.Cookie("ID_user")
	println(cookie, err)
	return cookie.Value
}

//envoie les donnees user recuperer par le cookies
func info_user(w http.ResponseWriter, req *http.Request) User {
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
	return data.User[profil(data, w, req)]
}

//---------------prendre l'index du mot associé---------------
func profil(data TabUser, w http.ResponseWriter, r *http.Request) int {
	var index2 int
	for i := 0; i < len(data.User); i++ {
		if data.User[i].ID_user == return_cookies(w, r) {
			index2 = i
		}
	}
	return index2
}

//fonction qui efface un cookie
func delete_cookies(w http.ResponseWriter, r *http.Request) {

	var cookie, err = r.Cookie("ID_user")
	fmt.Println(err)
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
}

//recup les info d'un id personne
func info_user_id(id string, w http.ResponseWriter, req *http.Request) User {
	dbUser := User{}
	infoUser, err := db.Query("SELECT id_user, id_role, firstname, lastname, pseudo, born, email, password, lang, creation FROM user WHERE id_user=?;", id)
	if err != nil {
		log.Fatal(err)
	}

	for infoUser.Next() {

		err := infoUser.Scan(&dbUser.ID_user, &dbUser.ID_role, &dbUser.Firstname, &dbUser.Lastname, &dbUser.Pseudo, &dbUser.Born, &dbUser.Email, &dbUser.Password, &dbUser.Lang, &dbUser.Creation)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(dbUser.ID_user, dbUser.ID_role, dbUser.Firstname, dbUser.Lastname, dbUser.Pseudo, dbUser.Born, dbUser.Email, dbUser.Password, dbUser.Lang)
	}
	defer infoUser.Close()
	return dbUser
}
