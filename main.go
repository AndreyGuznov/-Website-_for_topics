package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type New struct {
	Id      int
	Modek   string
	Company string
	Price   uint
}

var database *sql.DB //переменная для взаимодействия с БД

func IndexHandler(w http.ResponseWriter, r *http.Request) { //функция для отправки списка объектов из БД
	rows, err := database.Query("select * from new")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	new := []New{}
	for rows.Next() {
		n := New{}
		err := rows.Scan(&n.Id, &n.Modek, &n.Company, &n.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		new = append(new, n)
	}

	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, new)
}

func main() {

	db, err := sql.Open("mysql", "root:password@/newbase")
	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()
	http.HandleFunc("/about/", IndexHandler)

	fmt.Println("Server is listening..")
	http.ListenAndServe(":8080", nil)
}
