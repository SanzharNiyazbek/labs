//Sanzhar Niyazbek SIS-2122

package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Id     int
	Name   string
	Color  string
	Count  int
	Serial int
}

type server struct {
	db *sql.DB
}

func (s *server) warehousePage(w http.ResponseWriter, r *http.Request) {
	var products []Product
	res, _ := s.db.Query("SELECT * FROM warehouse;")
	for res.Next() {
		var product Product
		res.Scan(&product.Id, &product.Name, &product.Color, &product.Count, &product.Serial)
		products = append(products, product)
	}
	t, _ := template.ParseFiles("static/warehouse.html")
	t.Execute(w, products)
}

func (s *server) deletePage(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	_, err := s.db.Exec("DELETE FROM warehouse WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *server) updateProductPage(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("name") == "" {
		id := r.FormValue("id")
		Id := map[string]interface{}{"id": id}
		t, _ := template.ParseFiles("static/updateProduct.html")
		t.Execute(w, Id)
		return

	} else {
		id := r.FormValue("id")
		name := r.FormValue("name")
		color := r.FormValue("color")
		count := r.FormValue("count")
		serial := r.FormValue("serial")
		_, _ = s.db.Exec("UPDATE warehouse SET name=$1, color=$2, count=$3, serial=$4 WHERE id=$5", name, color, count, serial, id)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

func (s *server) addProductPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		color := r.FormValue("color")
		count := r.FormValue("count")
		serial := r.FormValue("serial")
		_, err := s.db.Exec("INSERT INTO warehouse(name, color, count, serial) VALUES ($1, $2, $3, $4)", name, color, count, serial)
		if err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/warehouse", http.StatusSeeOther)
	}
	t, _ := template.ParseFiles("static/addProduct.html")
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
	http.HandleFunc("/delete", s.deletePage)
	http.HandleFunc("/update", s.updateProductPage)
	http.HandleFunc("/addProduct", s.addProductPage)
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.ListenAndServe(":8000", nil)
}
