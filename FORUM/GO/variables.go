package main

import (
	"database/sql"
	"net/http"
)

// variables globales
var db *sql.DB
var w http.ResponseWriter
var r *http.Request
var mot string
var data2 TabForum
var email string
var password string
var err2 bool
