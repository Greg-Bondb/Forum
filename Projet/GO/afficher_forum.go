// package main

// import (
// 	"log"
// 	"net/http"
// 	"strings"
// 	"text/template"
// )

// //Structure message
// type Messages struct {
// 	ID_message   string
// 	Message      string
// 	Date_message string
// 	ID_forum     string
// 	ID_user      string
// }

// //Structure pour afficher un forum
// type InfoForum struct {
// 	ID_forum      string
// 	Title_subject string
// 	Categories    string
// 	Types         string
// 	Date          string
// 	ID_user       string
// 	Word          string
// 	Message       []Messages
// 	Error         bool
// }

// //Structure qui prend la liste de tous le forums
// type TabForum struct {
// 	ForumList []InfoForum
// 	Error     bool
// 	Error2    bool
// 	Pseudo    string
// }

// //Fonction qui regroupe deux routes (subject_affichage et forum+sujet)
// func route_affichage() {

// 	var data2 TabForum

// 	//---------------------------------
// 	http.HandleFunc("/subject_affichage", func(w http.ResponseWriter, r *http.Request) {
// 		tmpl4 := template.Must(template.ParseFiles("../HTML/affichage_forum.html"))

// 		//var forum a afficher
// 		data2 = TabForum{}
// 		forum := InfoForum{}

// 		if verif(w, r) {
// 			data2.Pseudo = info_user(w, r).Pseudo
// 			data2.Error = false
// 			data2.Error2 = false
// 		} else {
// 			data2.Pseudo = ""
// 			data2.Error = true
// 			data2.Error2 = err2
// 		}

// 		//afficher tous les forums
// 		rows, err := db.Query("SELECT tabforum.id_forum, tabforum.title_subject,tabforum.categories,tabforum.types,tabforum.messages,tabforum.date, tabforum.id_user FROM tabforum")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer rows.Close()

// 		for rows.Next() {

// 			err := rows.Scan(&forum.ID_forum, &forum.Title_subject, &forum.Categories, &forum.Types, &forum.Word, &forum.Date, &forum.ID_user)
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			//fonction qui remplace l'id par le pseudo
// 			if forum.ID_user != "" {
// 				forum.ID_user = info_user_id(forum.ID_user, w, r).Pseudo
// 			}
// 			if forum.ID_user == "" {
// 				forum.ID_user = "Compte supprimé !"
// 			}

// 			data2.ForumList = append(data2.ForumList, forum)
// 		}

// 		tmpl4.Execute(w, data2)
// 	})

// 	var mot string

// 	//---------------------------------
// 	//afficher un forum avec les messages
// 	http.HandleFunc("/forum/", func(w http.ResponseWriter, r *http.Request) {
// 		mot = strings.ReplaceAll(r.URL.Path, "/forum/", "")
// 		log.Println(mot)
// 		tmpl := template.Must(template.ParseFiles("../HTML/page_forum.html"))

// 		message := Messages{}

// 		//afficher tous les messages
// 		rows2, err2 := db.Query("SELECT id_message, message,date_message,id_forum,id_user FROM messages WHERE messages.id_forum = ?", data2.ForumList[forum(data2, mot)].ID_forum)
// 		if err2 != nil {
// 			log.Fatal(err2)
// 		}
// 		defer rows2.Close()

// 		for rows2.Next() {

// 			err2 := rows2.Scan(&message.ID_message, &message.Message, &message.Date_message, &message.ID_forum, &message.ID_user)
// 			if err2 != nil {
// 				log.Fatal(err2)
// 			}

// 			//fonction qui remplace l'id par le pseudo
// 			if message.ID_user != "" {
// 				message.ID_user = info_user_id(message.ID_user, w, r).Pseudo
// 			}
// 			if message.ID_user == "" {
// 				message.ID_user = "Compte supprimé !"
// 			}

// 			data2.ForumList[forum(data2, mot)].Message = append(data2.ForumList[forum(data2, mot)].Message, message)
// 			log.Println(message)
// 		}

// 		if verif(w, r) == true {
// 			data2.ForumList[forum(data2, mot)].Error = false
// 		} else {
// 			data2.ForumList[forum(data2, mot)].Error = true
// 		}

// 		tmpl.Execute(w, data2.ForumList[forum(data2, mot)])
// 		data2.ForumList[forum(data2, mot)].Message = nil
// 		log.Println(data2.ForumList[forum(data2, mot)].Message)
// 	})

// 	//---------------------------------
// 	//envoie de message
// 	http.HandleFunc("/message_exec", func(w http.ResponseWriter, r *http.Request) {

// 		if verif(w, r) == true && r.FormValue("message") != "" {
// 			envoie := Messages{
// 				Message:  r.FormValue("message"),
// 				ID_user:  return_cookies(w, r),
// 				ID_forum: data2.ForumList[forum(data2, mot)].ID_forum,
// 			}

// 			var err error
// 			_, err = db.Exec("INSERT INTO messages(id_message, message, date_message, id_forum, id_user) VALUES (DEFAULT,?,NOW(),?,?)", envoie.Message, envoie.ID_forum, envoie.ID_user)
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			http.Redirect(w, r, "/forum/"+mot, http.StatusFound)
// 		} else if verif(w, r) == true && r.FormValue("message") == "" {
// 			http.Redirect(w, r, "/forum/"+mot, http.StatusFound)
// 		} else if verif(w, r) == false {
// 			http.Redirect(w, r, "/forum/"+mot, http.StatusFound)
// 		}
// 	})

// 	//affichage des sujets populaires
// 	http.HandleFunc("/subject_pop", func(w http.ResponseWriter, r *http.Request) {
// 		tmpl4 := template.Must(template.ParseFiles("../HTML/affichage_forum.html"))

// 		//var forum a afficher
// 		data2 = TabForum{}

// 		if verif(w, r) {
// 			data2.Pseudo = info_user(w, r).Pseudo
// 			data2.Error = false
// 			data2.Error2 = false
// 		} else {
// 			data2.Pseudo = ""
// 			data2.Error = true
// 			data2.Error2 = err2
// 		}

// 		//afficher tous les forums populaires par ordre decroissant
// 		rows, err := db.Query("SELECT t.id_forum,t.title_subject,t.categories,t.types,t.date,t.id_user FROM messages AS m INNER JOIN tabforum AS t ON m.id_forum=T.id_forum GROUP BY m.id_forum ORDER BY COUNT(m.id_forum) DESC")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer rows.Close()

// 		for rows.Next() {

// 			forum := InfoForum{}

// 			err := rows.Scan(&forum.ID_forum, &forum.Title_subject, &forum.Categories, &forum.Types, &forum.Date, &forum.ID_user)
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			//fonction qui remplace l'id par le pseudo
// 			if forum.ID_user != "" {
// 				forum.ID_user = info_user_id(forum.ID_user, w, r).Pseudo
// 			}
// 			if forum.ID_user == "" {
// 				forum.ID_user = "Compte supprimé !"
// 			}

// 			data2.ForumList = append(data2.ForumList, forum)
// 		}

// 		tmpl4.Execute(w, data2)
// 	})

// }

// //---------------------------------
// //Fonction affichage reverse ACTU
// func route_actu() {
// 	var data2 TabForum

// 	//---------------------------------
// 	http.HandleFunc("/subject_actu", func(w http.ResponseWriter, r *http.Request) {
// 		tmpl4 := template.Must(template.ParseFiles("../HTML/affichage_forum.html"))

// 		data2 = TabForum{}

// 		if verif(w, r) {
// 			data2.Pseudo = info_user(w, r).Pseudo
// 			data2.Error = false
// 			data2.Error2 = false
// 		} else {
// 			data2.Pseudo = ""
// 			data2.Error = true
// 			data2.Error2 = err2
// 		}

// 		rows, err := db.Query("SELECT id_user, title_subject, categories, types, date FROM tabforum;")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer rows.Close()
// 		for rows.Next() {

// 			forum := InfoForum{}

// 			err := rows.Scan(&forum.ID_user, &forum.Title_subject, &forum.Categories, &forum.Types, &forum.Date)
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			//fonction qui remplace l'id par le pseudo
// 			if forum.ID_user != "" {
// 				forum.ID_user = info_user_id(forum.ID_user, w, r).Pseudo
// 			}
// 			if forum.ID_user == "" {
// 				forum.ID_user = "Compte supprimé !"
// 			}

// 			data2.ForumList = append(data2.ForumList, forum)
// 		}
// 		for i, j := 0, len(data2.ForumList)-1; i < j; i, j = i+1, j-1 {
// 			data2.ForumList[i], data2.ForumList[j] = data2.ForumList[j], data2.ForumList[i]
// 		}
// 		tmpl4.Execute(w, data2)
// 	})

// }

// //---------------prendre l'index du mot associÃ©---------------
// func forum(data TabForum, mot string) int {
// 	var index int
// 	for i := 0; i < len(data.ForumList); i++ {
// 		if data.ForumList[i].Title_subject == string(mot) {
// 			index = i
// 		}
// 	}
// 	return index
// }

package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

//Structure message
type Messages struct {
	ID_message   string
	Message      string
	Date_message string
	ID_forum     string
	ID_user      string
}

//Structure pour afficher un forum
type InfoForum struct {
	ID_forum      string
	Title_subject string
	Categories    string
	Types         string
	Date          string
	ID_user       string
	Word          string
	Message       []Messages
	Error         bool
	Error1        bool
	Error2        bool
	Pseudo        string
}

//Structure qui prend la liste de tous le forums
type TabForum struct {
	ForumList []InfoForum
	Error     bool
	Error2    bool
	Pseudo    string
}

//Fonction qui regroupe deux routes (subject_affichage et forum+sujet)
func route_affichage() {

	var data2 TabForum

	//---------------------------------
	http.HandleFunc("/subject_affichage", func(w http.ResponseWriter, r *http.Request) {
		tmpl4 := template.Must(template.ParseFiles("../HTML/affichage_forum.html"))

		//var forum a afficher
		data2 = TabForum{}
		forum := InfoForum{}

		if verif(w, r) {
			data2.Pseudo = info_user(w, r).Pseudo
			data2.Error = false
			data2.Error2 = false
		} else {
			data2.Pseudo = ""
			data2.Error = true
			data2.Error2 = err2
		}

		//afficher tous les forums
		rows, err := db.Query("SELECT tabforum.id_forum, tabforum.title_subject,tabforum.categories,tabforum.types,tabforum.messages,tabforum.date, tabforum.id_user FROM tabforum")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {

			err := rows.Scan(&forum.ID_forum, &forum.Title_subject, &forum.Categories, &forum.Types, &forum.Word, &forum.Date, &forum.ID_user)
			if err != nil {
				log.Fatal(err)
			}

			//fonction qui remplace l'id par le pseudo
			if forum.ID_user != "" {
				forum.ID_user = info_user_id(forum.ID_user, w, r).Pseudo
			}
			if forum.ID_user == "" {
				forum.ID_user = "Compte supprimé !"
			}

			data2.ForumList = append(data2.ForumList, forum)
		}

		tmpl4.Execute(w, data2)
	})

	var mot string

	//---------------------------------
	//afficher un forum avec les messages
	http.HandleFunc("/forum/", func(w http.ResponseWriter, r *http.Request) {
		mot = strings.ReplaceAll(r.URL.Path, "/forum/", "")
		log.Println(mot)
		tmpl := template.Must(template.ParseFiles("../HTML/page_forum.html"))

		message := Messages{}
		infoforum := InfoForum{}

		//afficher tous les messages
		rows2, err2 := db.Query("SELECT id_message, message,date_message,id_forum,id_user FROM messages WHERE messages.id_forum = ?", data2.ForumList[forum(data2, mot)].ID_forum)
		if err2 != nil {
			log.Fatal(err2)
		}
		defer rows2.Close()

		for rows2.Next() {

			err2 := rows2.Scan(&message.ID_message, &message.Message, &message.Date_message, &message.ID_forum, &message.ID_user)
			if err2 != nil {
				log.Fatal(err2)
			}

			//fonction qui remplace l'id par le pseudo
			if message.ID_user != "" {
				message.ID_user = info_user_id(message.ID_user, w, r).Pseudo
				infoforum.Error1 = false
				infoforum.Error2 = false
			}
			if message.ID_user == "" {
				message.ID_user = "Compte supprimé !"
			}

			data2.ForumList[forum(data2, mot)].Message = append(data2.ForumList[forum(data2, mot)].Message, message)
			log.Println(message)
		}

		if verif(w, r) == true {
			data2.ForumList[forum(data2, mot)].Error = false
		} else {
			data2.ForumList[forum(data2, mot)].Error = true
		}

		tmpl.Execute(w, data2.ForumList[forum(data2, mot)])
		data2.ForumList[forum(data2, mot)].Message = nil
		log.Println(data2.ForumList[forum(data2, mot)].Message)
	})

	//---------------------------------
	//envoie de message
	http.HandleFunc("/message_exec", func(w http.ResponseWriter, r *http.Request) {

		if verif(w, r) == true && r.FormValue("message") != "" {
			envoie := Messages{
				Message:  r.FormValue("message"),
				ID_user:  return_cookies(w, r),
				ID_forum: data2.ForumList[forum(data2, mot)].ID_forum,
			}

			var err error
			_, err = db.Exec("INSERT INTO messages(id_message, message, date_message, id_forum, id_user) VALUES (DEFAULT,?,NOW(),?,?)", envoie.Message, envoie.ID_forum, envoie.ID_user)
			if err != nil {
				log.Fatal(err)
			}

			http.Redirect(w, r, "/forum/"+mot, http.StatusFound)
		} else if verif(w, r) == true && r.FormValue("message") == "" {
			http.Redirect(w, r, "/forum/"+mot, http.StatusFound)
		} else if verif(w, r) == false {
			http.Redirect(w, r, "/forum/"+mot, http.StatusFound)
		}
	})

	//affichage des sujets populaires
	http.HandleFunc("/subject_pop", func(w http.ResponseWriter, r *http.Request) {
		tmpl4 := template.Must(template.ParseFiles("../HTML/affichage_forum.html"))

		//var forum a afficher
		data2 = TabForum{}

		if verif(w, r) {
			data2.Pseudo = info_user(w, r).Pseudo
			data2.Error = false
			data2.Error2 = false
		} else {
			data2.Pseudo = ""
			data2.Error = true
			data2.Error2 = err2
		}

		//afficher tous les forums populaires par ordre decroissant
		rows, err := db.Query("SELECT t.id_forum,t.title_subject,t.categories,t.types,t.date,t.id_user FROM messages AS m INNER JOIN tabforum AS t ON m.id_forum=T.id_forum GROUP BY m.id_forum ORDER BY COUNT(m.id_forum) DESC")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {

			forum := InfoForum{}

			err := rows.Scan(&forum.ID_forum, &forum.Title_subject, &forum.Categories, &forum.Types, &forum.Date, &forum.ID_user)
			if err != nil {
				log.Fatal(err)
			}

			//fonction qui remplace l'id par le pseudo
			if forum.ID_user != "" {
				forum.ID_user = info_user_id(forum.ID_user, w, r).Pseudo
			}
			if forum.ID_user == "" {
				forum.ID_user = "Compte supprimé !"
			}

			data2.ForumList = append(data2.ForumList, forum)
		}

		tmpl4.Execute(w, data2)
	})

}

//---------------------------------
//Fonction affichage reverse ACTU
func route_actu() {
	var data2 TabForum

	//---------------------------------
	http.HandleFunc("/subject_actu", func(w http.ResponseWriter, r *http.Request) {
		tmpl4 := template.Must(template.ParseFiles("../HTML/affichage_forum.html"))

		data2 = TabForum{}

		if verif(w, r) {
			data2.Pseudo = info_user(w, r).Pseudo
			data2.Error = false
			data2.Error2 = false
		} else {
			data2.Pseudo = ""
			data2.Error = true
			data2.Error2 = err2
		}

		rows, err := db.Query("SELECT id_user, title_subject, categories, types, date FROM tabforum;")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {

			forum := InfoForum{}

			err := rows.Scan(&forum.ID_user, &forum.Title_subject, &forum.Categories, &forum.Types, &forum.Date)
			if err != nil {
				log.Fatal(err)
			}

			//fonction qui remplace l'id par le pseudo
			if forum.ID_user != "" {
				forum.ID_user = info_user_id(forum.ID_user, w, r).Pseudo
			}
			if forum.ID_user == "" {
				forum.ID_user = "Compte supprimé !"
			}

			data2.ForumList = append(data2.ForumList, forum)
		}
		for i, j := 0, len(data2.ForumList)-1; i < j; i, j = i+1, j-1 {
			data2.ForumList[i], data2.ForumList[j] = data2.ForumList[j], data2.ForumList[i]
		}
		tmpl4.Execute(w, data2)
	})

}

//---------------prendre l'index du mot associÃ©---------------
func forum(data TabForum, mot string) int {
	var index int
	for i := 0; i < len(data.ForumList); i++ {
		if data.ForumList[i].Title_subject == string(mot) {
			index = i
		}
	}
	return index
}
