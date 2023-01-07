package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func handle_supprime_message(w http.ResponseWriter, r *http.Request) {
	if verif(w, r) == true {
		if info_user(w, r).ID_role == "1" {

			id_message := r.FormValue("delete")
			intVar, err2 := strconv.Atoi(id_message)
			log.Println(intVar, err2)

			var err error
			log.Println(id_message)

			_, err = db.Exec(`DELETE FROM messages WHERE id_message=? ;`, intVar)
			if err != nil {
				panic(err)
			}

			fmt.Println("Le profil a bien été supprimé !")

			http.Redirect(w, r, "/accueil", http.StatusFound)
		}
	} else {
		http.Redirect(w, r, "/accueil", http.StatusFound)
	}
}

func handle_supprime_forum(w http.ResponseWriter, r *http.Request) {
	if verif(w, r) == true {
		if info_user(w, r).ID_role == "1" {
			id_forum := r.FormValue("delete")
			intVar, err2 := strconv.Atoi(id_forum)
			log.Println(intVar, err2)

			var err error
			log.Println(id_forum)
			_, err = db.Exec(`DELETE FROM tabforum WHERE id_forum=? ;`, intVar)
			if err != nil {
				panic(err)
			}

			fmt.Println("Le profil a bien été supprimé !")

			http.Redirect(w, r, "/accueil", http.StatusFound)
		}
	} else {
		http.Redirect(w, r, "/accueil", http.StatusFound)
	}
}
