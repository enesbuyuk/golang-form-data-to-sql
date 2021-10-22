package main

import (
	"html/template"
	"database/sql"
	"net/http"
	"log"
	_ "github.com/go-sql-driver/mysql"

)

var yapi *template.Template

func main() {

	dbhost := "localhost"
	dbport := "3306"
	dbname := "enesbuyukcom"
	dbuser := "root"
	dbpass := "root"

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp("+dbhost+":"+dbport+")/"+dbname)
	if err != nil {
		log.Print(err.Error())
	}


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
	    	if(len(r.FormValue("title")) != 0){
				ekle, err := db.Prepare(`INSERT INTO example (title,content) VALUES (?,?);`)
				if err != nil {
					panic(err.Error())
				}
				ekle.Exec(r.FormValue("title"),r.FormValue("content"))
	  			tmpl := template.Must(template.ParseFiles("success.html"))
				tmpl.Execute(w, nil)
	    	}
	 	}else{
	 		tmpl := template.Must(template.ParseFiles("index.html"))
  			 tmpl.Execute(w, nil)
		}
	})
	http.ListenAndServe(":80", nil)
	
}
