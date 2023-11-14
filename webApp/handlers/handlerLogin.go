package forum

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var data Login
	data.User_email = r.PostFormValue("user_email2")
	if !isUsernameExists(db, data.User_email) {
		fmt.Println("User Email not exist!")
	}
	data.User_pass = r.PostFormValue("password2")
	rows, err := db.Query("SELECT user_pass FROM users WHERE user_email = ?", data.User_email)
	if err != nil {
		fmt.Println(err)
	}
	var current_pass string
	for rows.Next() {
		rows.Scan(&current_pass)
		break
	}
	if data.User_pass == current_pass {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	t, err := template.ParseFiles(HTMLs...)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "test.html", nil)
}

func isUsernameExists(db *sql.DB, user_email string) bool {
	sqlStmt, err := db.Prepare("SELECT COUNT(*) FROM users WHERE user_email = ?")
	if err != nil {
		return false
	}
	defer sqlStmt.Close()

	var count int
	err = sqlStmt.QueryRow(user_email).Scan(&count)
	if err != nil {
		return false
	}

	return count > 0
}
