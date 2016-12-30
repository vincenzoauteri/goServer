package main

import (
    "fmt"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "html/template"
    "net/http"
    "regexp"
    "os"
)

type Page struct {
    Title string
    Body  []byte
}

var staticPath = "static\\templates\\"
var templates = template.Must(template.ParseFiles(staticPath+"login.html",
    staticPath+"register.html",
    staticPath+"forgottenPassword.html",
    staticPath+"index.html"))

var validPath = regexp.MustCompile("^/(login)/([a-zA-Z0-9]+)$")

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html",p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
    p := Page{Title: "Home", Body: nil}
    renderTemplate(w, "index" ,&p)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    p := Page{Title: "Login", Body: nil}
    renderTemplate(w, "login" ,&p)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    p := Page{Title: "Register", Body: nil}
    renderTemplate(w, "register" ,&p)
}

func forgottenPasswordHandler(w http.ResponseWriter, r *http.Request) {
    p := Page{Title: "Forgotten Password", Body: nil}
    renderTemplate(w, "forgottenPassword" ,&p)
}

func createTable() {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table users(id integer not null primary key, name text, password text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func initDB() {
    os.Remove("./users.db")
    createTable()
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
/*
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("test", i))
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()

	rows, err := db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("select name from foo where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	_, err = db.Exec("delete from foo")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	if err != nil {
		log.Fatal(err)
	}

	rows, err = db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
*/

func main() {
    fmt.Printf("Setting Up DB \n")
    initDB()

    fmt.Printf("Starting Web Server \n")
    http.HandleFunc("/", mainHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/register", registerHandler)
    http.HandleFunc("/forgottenPassword", forgottenPasswordHandler)
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    http.ListenAndServe(":8080", nil)
}
