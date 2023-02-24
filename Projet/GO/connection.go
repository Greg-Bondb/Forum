package main

import "net/http"

//---------------connexion dossier---------------
func connection() {
	cssglo := http.FileServer(http.Dir("../CSS"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", cssglo))

	img := http.FileServer(http.Dir("../Picture"))
	http.Handle("/Picture/", http.StripPrefix("/Picture/", img))

	js := http.FileServer(http.Dir("../JS"))
	http.Handle("/JS/", http.StripPrefix("/JS/", js))
}
