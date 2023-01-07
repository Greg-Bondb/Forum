package main

import (
	"fmt"
	"net/http"
)

//---------------récup suppr profil---------------
func handler_delete_exec(w http.ResponseWriter, r *http.Request) {
	if verif(w, r) == true {
		delete_profil := info_user(w, r)

		var err error

		sqlStatement := `DELETE FROM user WHERE email=? ;`
		_, err = db.Exec(sqlStatement, delete_profil.Email)
		if err != nil {
			panic(err)
		}

		delete_cookies(w, r)

		fmt.Println("Le profil a bien été supprimé !")

		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

}
