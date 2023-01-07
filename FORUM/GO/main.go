package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//---------------affichage register---------------

//---------------func main---------------
func main() {
	connection()
	connectDB()

	//---------------TEST---------------
	http.HandleFunc("/accueil", handler_login)
	http.HandleFunc("/login_exec", handler_login_exec)

	//---------------GOOD---------------
	creauser()

	//---------------GOOD---------------
	http.HandleFunc("/profil", handler_profil)

	//---------------GOOD---------------
	http.HandleFunc("/modifyprofil", handler_modifyprofil)
	http.HandleFunc("/modifyprofil_exec", handler_modifyprofil_exec)

	//---------------GOOD---------------
	http.HandleFunc("/delete_exec", handler_delete_exec)

	//---------------GOOD---------------
	http.HandleFunc("/subject", handler_creationforum)
	http.HandleFunc("/subject_exec", handler_creationforum_exec)

	//---------------GOOD---------------
	route_affichage()
	route_actu()

	http.HandleFunc("/delete_forum", handle_supprime_forum)
	http.HandleFunc("/delete_message", handle_supprime_message)
	//---------------GOOD---------------
	http.HandleFunc("/contact", handler_contact)

	//--------------TEST----------------
	http.HandleFunc("/security", handler_modifypassword)
	http.HandleFunc("/security_exec", handler_modifypassword_exec)

	http.HandleFunc("/login_delete", handler_login_delete)

	http.HandleFunc("/reglement", handler_reglement)
	http.HandleFunc("/politique", handler_politique)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
