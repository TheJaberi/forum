package forum

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sql.DB
)

func init() {
	var err error
	DB, err = sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	DataBase(DB)
	row, _ := DB.Query("SELECT user_id, user_name FROM users WHERE user_id = 1")
	var user_id int
	var user_name string
	for row.Next() {
		row.Scan(&user_id, &user_name)
		break
	}
	if user_id != 1 && user_name != "admin" {
		sqlStmt, _ := DB.Prepare("INSERT INTO users (user_name, user_email, user_pass, user_type) VALUES (?, ?, ?, ?)")
		sqlStmt.Exec("admin", "admin", "admin", "admin")
	}
}

func PathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/": // home page
		MainHandler(w, r)
	case "/sign_up": // artist page
		SignUpHandler(w, r, DB)
	// case "/sign_in": // about page
	// 	AboutHandler(w, r)
	// case "/create_post": // welcome page
	// 	WelcomeHundler(w, r)
	// case "/create_comment": // filter page
	// 	FilterHundler(w, r, data)
	// case "/post": // search page
	// 	SearchHundler(w, r, data)
	case "/400":
		ErrorHandler(w, r, 400)
	case "/500":
		ErrorHandler(w, r, 500)
	default: // any invalid path will be redirected to 404 error handler
		ErrorHandler(w, r, 404)
	}
}
