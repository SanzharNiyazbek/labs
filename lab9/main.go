//Sanzhar Niyazbek SIS-2122

package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type server struct {
	db *sql.DB
}

func (s *server) warehousePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		color := r.FormValue("color")
		count := r.FormValue("count")
		serial := r.FormValue("serial")
		_, err := s.db.Exec("INSERT INTO warehouse(name, color, count, serial) VALUES ($1, $2, $3, $4)", name, color, count, serial)
		if err != nil {
			log.Fatal(err)
		}
		message := map[string]interface{}{"msg": "Success!"}
		t, _ := template.ParseFiles("static/warehouse.html")
		t.Execute(w, message)
		return
	}
	t, _ := template.ParseFiles("static/warehouse.html")
	t.Execute(w, nil)
}

func dbConnect() server { // Server connection
	db, _ := sql.Open("sqlite3", "data.db")
	s := server{db: db}
	return s
}

func main() { //main function
	s := dbConnect()
	defer s.db.Close()
	http.HandleFunc("/warehouse", s.warehousePage)
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.ListenAndServe(":8000", nil)
}
