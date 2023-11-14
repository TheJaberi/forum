package forum

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	bcrypt "golang.org/x/crypto/bcrypt"
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
	sqlStmt, err := db.Prepare("SELECT user_pass FROM users WHERE user_email = ?")
	if err != nil {
		fmt.Println(err)
	}
	var session Session = Session{}
	var current_pass []byte
	err = sqlStmt.QueryRow(data.User_email).Scan(&current_pass)
	if err != nil {
		fmt.Println("Error: in reading password or email not exist!")
	} else {
		fmt.Println("Your pass match the password in data!")
		session.Email = data.User_email
		session.Uuid, _ = uuid.NewV4()
		session.CreatedAt = time.Now()
		fmt.Println(session)
	}
	err = bcrypt.CompareHashAndPassword(current_pass, []byte(data.User_pass))
	if err != nil {
		fmt.Println("Error: not match!")
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
