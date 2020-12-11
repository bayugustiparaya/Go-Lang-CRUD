package main

import (
	"database/sql"
	// "encoding/json"
	// "fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

// Mahasiswa struct (Model)
type Mahasiswa struct {
	Nobp		string	`json:"nobp"`
	Nama		string	`json:"nama"`
	Alamat	string	`json:"alamat"`
	Umur		int			`json:"umur"`
}

func main() {
	db, err = sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/go-cruds")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	r := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":9000", r))
}