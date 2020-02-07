package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	db, dbErr := sql.Open("mysql", "root:@/learning_go?charset=utf8mb4")
	if dbErr != nil {
		log.Fatal("DB Error: ", dbErr)
	}
	stmt, dbErr := db.Prepare("INSERT users SET display_name=?,created_at=?")
	if dbErr != nil {
		log.Fatal("DB Prepare Error: ", dbErr)
	}
	res, dbErr := stmt.Exec("Test User", "2020-02-02")
	if dbErr != nil {
		log.Fatal("DB Exec Error: ", dbErr)
	}
	id, dbErr := res.LastInsertId()
	if dbErr != nil {
		log.Fatal("DB LastInsertId Error: ", dbErr)
	}
	log.Fatal("DB Id: ", id)

	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
