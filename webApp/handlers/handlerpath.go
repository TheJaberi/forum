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
	if !isTableExists(DB, "users") {
		DataBase(DB)
		sqlStmt, err := DB.Prepare("INSERT INTO users (user_name, user_email, user_pass, user_type) VALUES (?, ?, ?, ?)")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sqlStmt.Exec("admin", "admin", "admin", "admin")
	}

}

func PathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/": // home page
		MainHandler(w, r)
	case "/sign_up": // sign up page
		SignUpHandler(w, r, DB)
	case "/sign_in": // sign in page
		SignInHandler(w, r, DB)
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

func isTableExists(db *sql.DB, tableName string) bool {
	sqlStmt, err := db.Prepare("SELECT COUNT(*) FROM sqlite_master WHERE type = 'table' AND name = ?")
	if err != nil {
		return false
	}
	defer sqlStmt.Close()
	var count int
	err = sqlStmt.QueryRow(tableName).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}
